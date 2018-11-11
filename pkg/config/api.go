package config

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

type Config interface {
	StorePullRequest(remote string, request *PullRequest) error
	Clean() error
	GetRemoteURL(remote string) (string, error)
	ListRemoteNames() ([]string, error)
	GetCurrentBranch() (string, error)
	GetPullRequestForBranch(branch string) (*PullRequest, error)
}

func NewGitConfig() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(wd)
	if err != nil {
		return nil, err
	}

	return &gitConfig{repo}, nil
}
