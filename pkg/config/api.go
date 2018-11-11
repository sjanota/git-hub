package config

import (
	"github.com/google/go-github/v18/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

type Config interface {
	StorePullRequest(remote string, request *github.PullRequest) error
	Clean() error
	GetRemoteURL(remote string) (string, error)
	ListRemoteNames() ([]string, error)
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
