package config

import (
	"fmt"
	"github.com/google/go-github/v18/github"
	"gopkg.in/src-d/go-git.v4"
	"strconv"
)

type gitConfig struct {
	repo *git.Repository
}

func (c *gitConfig) ListRemoteNames() ([]string, error) {
	remotes, err := c.repo.Remotes()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(remotes))
	for i, remote := range remotes {
		result[i] = remote.Config().Name
	}
	return result, nil
}

func (c *gitConfig) GetRemoteURL(remoteName string) (string, error) {
	remote, err := c.repo.Remote(remoteName)
	if err != nil {
		return "", err
	}

	return remote.Config().URLs[0], nil
}

func (c *gitConfig) Clean() error {
	config, err := c.repo.Config()
	if err != nil {
		return err
	}

	config.Raw.RemoveSection("github-pr")

	err = c.repo.Storer.SetConfig(config)
	if err != nil {
		return err
	}

	return nil
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
