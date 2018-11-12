package git

import (
	"fmt"
	"github.com/pkg/errors"
	git_config "gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type PullRequest struct {
	HeadRef  string
	HeadRepo string
	Number   int
	WebURL   string
	Remote   string
	Title    string
	Comment  string
	InSync   bool
}

const (
	pullRequestSection  = "github-pr"
	pullRequestTitle    = "title"
	pullRequestComment  = "comment"
	pullRequestHeadRef  = "headRef"
	pullRequestHeadRepo = "headRepo"
	pullRequestWebUrl   = "webUrl"
	pullRequestInSync   = "inSync"
)

func (r *repository) ListPullRequests() ([]*PullRequest, error) {
	cfg, err := r.repo.Config()
	if err != nil {
		return nil, err
	}

	subsections := cfg.Raw.Section("github-pr").Subsections
	result := make([]*PullRequest, len(subsections))

	for i, subsection := range subsections {
		result[i], err = readPullRequestFromSubsection(subsection)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (r *repository) StorePullRequest(pr *PullRequest) error {
	cfg, err := r.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", pr.Remote, pr.Number)

	cfg.Raw.Section(pullRequestSection).Subsection(subsection).
		SetOption(pullRequestTitle, pr.Title).
		SetOption(pullRequestComment, strings.Replace(pr.Comment, "\n", "\\n", -1)).
		SetOption(pullRequestHeadRef, pr.HeadRef).
		SetOption(pullRequestHeadRepo, pr.HeadRepo).
		SetOption(pullRequestWebUrl, pr.WebURL).
		SetOption(pullRequestInSync, fmt.Sprintf("%v", pr.InSync))

	err = r.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPullRequestForBranch(branch string) (*PullRequest, error) {
	cfg, err := r.repo.Config()
	if err != nil {
		return nil, err
	}

	for _, subsection := range cfg.Raw.Section("github-pr").Subsections {
		if subsection.Option("headRef") == branch {
			return readPullRequestFromSubsection(subsection)
		}
	}
	return nil, NoPullRequestForBranch{Branch: branch}
}

func readPullRequestFromSubsection(subsection *git_config.Subsection) (*PullRequest, error) {
	name := strings.Split(subsection.Name, ":")
	number, err := strconv.Atoi(name[1])
	if err != nil {
		return nil, err
	}

	inSync, err := strconv.ParseBool(subsection.Option(pullRequestInSync))
	if err != nil {
		return &PullRequest{}, err
	}

	return &PullRequest{
		Title:    subsection.Option(pullRequestTitle),
		Comment:  strings.Replace(subsection.Option(pullRequestComment), "\\n", "\n", -1),
		HeadRef:  subsection.Option(pullRequestHeadRef),
		HeadRepo: subsection.Option(pullRequestHeadRepo),
		Number:   number,
		WebURL:   subsection.Option(pullRequestWebUrl),
		Remote:   name[0],
		InSync:   inSync,
	}, nil
}

type CommentEditor interface {
	Edit(pr *PullRequest) (new string, err error)
}

type staticCommentEditor struct {
	comment string
	append  bool
}

func (r repository) StaticCommentEditor(comment string, append bool) CommentEditor {
	return staticCommentEditor{comment: comment, append: append}
}

func (e staticCommentEditor) Edit(pr *PullRequest) (string, error) {
	if e.append {
		return fmt.Sprintf("%s\n%s", pr.Comment, e.comment), nil
	}
	return e.comment, nil
}

func (r repository) FileCommentEditor() CommentEditor {
	return fileCommentEditor{repo: r}
}

type fileCommentEditor struct {
	repo repository
}

func (e fileCommentEditor) Edit(pr *PullRequest) (string, error) {
	rootDir, err := e.repo.GetRootDir()
	if err != nil {
		return "", err
	}

	gitDir := filepath.Join(rootDir, ".git")
	fileName := fmt.Sprintf("PR_%v_COMMENT", pr.Number)
	f, err := ioutil.TempFile(gitDir, fileName)
	if err != nil {
		return "", errors.Wrapf(err, "cannot create temp file %s", fileName)
	}
	defer os.Remove(f.Name())

	_, err = f.WriteString(pr.Comment)
	if err != nil {
		return "", errors.Wrap(err, "cannot write comment to file")
	}

	editor, err := e.repo.GetDefaultTextEditor()
	if err != nil {
		return "", errors.Wrap(err, "cannot get text editor")
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(editor, f.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", errors.Wrapf(err, "cannot open text editor %s", editor)
	}

	output, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return "", errors.Wrap(err, "cannot read comment content")
	}

	return strings.TrimSpace(string(output)), nil
}
