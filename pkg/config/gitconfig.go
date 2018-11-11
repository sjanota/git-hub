package config

import (
	"fmt"
	"github.com/google/go-github/v18/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"strconv"
)

type Config interface {
	StorePullRequest(remote string, request *github.PullRequest) error
}

type gitConfig struct {
	repo *git.Repository
}

func (c *gitConfig) StorePullRequest(remote string, request *github.PullRequest) error {
	config, err := c.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", remote, request.GetNumber())

	config.Raw.Section("github-pr").Subsection(subsection).
		SetOption("headRef", *request.Head.Ref).
		SetOption("headRepo", *request.Head.Repo.FullName).
		SetOption("baseRef", *request.Base.Ref).
		SetOption("baseRepo", *request.Base.Repo.FullName).
		SetOption("number", strconv.Itoa(*request.Number))

	err = c.repo.Storer.SetConfig(config)
	if err != nil {
		return err
	}

	return nil
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
