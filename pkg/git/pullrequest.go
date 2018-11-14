package git

import (
	"encoding/base64"
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
	State    string
}

const (
	prSection  = "github-pr"
	prTitle    = "title"
	prComment  = "comment"
	prHeadRef  = "headRef"
	prHeadRepo = "headRepo"
	prWebUrl   = "webUrl"
	prState    = "state"
)

func (r *repository) RemovePR(remote string, number int) error {
	cfg, err := r.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", remote, number)

	cfg.Raw.RemoveSubsection("github-pr", subsection)

	return r.repo.Storer.SetConfig(cfg)
}

func (r *repository) ListPRs() ([]*PullRequest, error) {
	cfg, err := r.repo.Config()
	if err != nil {
		return nil, err
	}

	subsections := cfg.Raw.Section("github-pr").Subsections
	result := make([]*PullRequest, len(subsections))

	for i, subsection := range subsections {
		result[i], err = readPRFromSubsection(subsection)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (r *repository) StorePR(pr *PullRequest) error {
	cfg, err := r.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", pr.Remote, pr.Number)

	cfg.Raw.Section(prSection).Subsection(subsection).
		SetOption(prTitle, pr.Title).
		SetOption(prComment, encodeComment(pr.Comment)).
		SetOption(prHeadRef, pr.HeadRef).
		SetOption(prHeadRepo, pr.HeadRepo).
		SetOption(prState, pr.State).
		SetOption(prWebUrl, pr.WebURL)

	err = r.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPR(remote string, number int) (*PullRequest, error) {
	cfg, err := r.repo.Config()
	if err != nil {
		return nil, err
	}

	subsection := fmt.Sprintf("%s:%v", remote, number)

	if !cfg.Raw.Section(prSection).HasSubsection(subsection) {
		return nil, PRNotFound{Remote: remote, Number: number}
	}

	return readPRFromSubsection(cfg.Raw.Section(prSection).Subsection(subsection))
}

func (r *repository) GetPRForBranch(branch string) (*PullRequest, error) {
	cfg, err := r.repo.Config()
	if err != nil {
		return nil, err
	}

	for _, subsection := range cfg.Raw.Section("github-pr").Subsections {
		if subsection.Option("headRef") == branch {
			return readPRFromSubsection(subsection)
		}
	}
	return nil, NoPRForBranch{Branch: branch}
}

func readPRFromSubsection(subsection *git_config.Subsection) (*PullRequest, error) {
	name := strings.Split(subsection.Name, ":")
	number, err := strconv.Atoi(name[1])
	if err != nil {
		return nil, err
	}

	decodedComment, err := decodeComment(subsection.Option(prComment))
	if err != nil {
		return nil, err
	}

	return &PullRequest{
		Title:    subsection.Option(prTitle),
		Comment:  decodedComment,
		HeadRef:  subsection.Option(prHeadRef),
		State:    subsection.Option(prState),
		HeadRepo: subsection.Option(prHeadRepo),
		Number:   number,
		WebURL:   subsection.Option(prWebUrl),
		Remote:   name[0],
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
	defer func() { _ = os.Remove(f.Name()) }()

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

func encodeComment(decoded string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(decoded))
}

func decodeComment(encoded string) (string, error) {
	bytes, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
