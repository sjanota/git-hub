package git

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

type Config interface {
	StorePullRequest(request *PullRequest) error
	Clean() error
	GetRemoteURL(remote string) (string, error)
	ListRemoteNames() ([]string, error)
	GetCurrentBranch() (string, error)
	GetPullRequestForBranch(branch string) (*PullRequest, error)
	ListPullRequests() ([]*PullRequest, error)
	GetDefaultTextEditor() (string, error)
}

func NewConfig() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(wd)
	if err != nil {
		return nil, err
	}

	return &config{repo}, nil
}
