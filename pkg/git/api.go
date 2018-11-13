package git

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
)

type Repo interface {
	StorePullRequest(pr *PullRequest) error
	ListPullRequests() ([]*PullRequest, error)
	GetPullRequest(remote string, number int) (*PullRequest, error)
	GetPullRequestForBranch(branch string) (*PullRequest, error)

	StaticCommentEditor(comment string, append bool) CommentEditor
	FileCommentEditor() CommentEditor

	GetRemoteURL(remote string) (string, error)
	GetRemoteForURL(url string) (string, error)
	ListRemoteNames() ([]string, error)

	GetCurrentBranch() (string, error)
	GetBranchRemote(branch string) (string, error)

	GetDefaultTextEditor() (string, error)
	GetRootDir() (string, error)
	GetCredentials(remote string) (*Credentials, error)
	Clean() error
}

func NewConfig() (Repo, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rootDir, err := findRootDir(wd)
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(rootDir)
	if err != nil {
		return nil, err
	}

	return &repository{repo}, nil
}

func findRootDir(wd string) (string, error) {
	tmpWd := wd
	for {
		gitDir := filepath.Join(tmpWd, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return tmpWd, nil
		} else if err != nil && !os.IsNotExist(err) {
			return "", err
		}

		newWd := filepath.Dir(tmpWd)
		if newWd == tmpWd {
			return "", errors.Errorf("not git directory %s", wd)
		}
		tmpWd = newWd
	}
}
