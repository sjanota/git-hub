package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"strconv"
	"strings"
)

type gitConfig struct {
	repo *git.Repository
}

func (c *gitConfig) GetPullRequestForBranch(branch string) (*PullRequest, error) {
	cfg, err := c.repo.Config()
	if err != nil {
		return nil, err
	}

	for _, subsection := range cfg.Raw.Section("github-pr").Subsections {
		if subsection.Option("headRef") == branch {
			return readPullRequestFromSubsection(subsection)
		}
	}
	return nil, errors.Errorf("no PR for branch %s", branch)
}

func readPullRequestFromSubsection(subsection *config.Subsection) (*PullRequest, error) {
	name := strings.Split(subsection.Name, ":")
	number, err := strconv.Atoi(name[1])
	if err != nil {
		return nil, err
	}
	return &PullRequest{
		HeadRef:  subsection.Option("headRef"),
		HeadRepo: subsection.Option("headRepo"),
		BaseRef:  subsection.Option("baseRef"),
		BaseRepo: subsection.Option("baseRepo"),
		Number:   number,
		WebURL:   subsection.Option("webUrl"),
		Remote:   name[0],
	}, nil
}

func (c *gitConfig) GetCurrentBranch() (string, error) {
	ref, err := c.repo.Head()
	if err != nil {
		return "", err
	}

	if !ref.Name().IsBranch() {
		return "", errors.Errorf("detached head")
	}

	return ref.Name().Short(), nil
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
	cfg, err := c.repo.Config()
	if err != nil {
		return err
	}

	cfg.Raw.RemoveSection("github-pr")

	err = c.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (c *gitConfig) StorePullRequest(remote string, request *PullRequest) error {
	cfg, err := c.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", remote, request.Number)

	cfg.Raw.Section("github-pr").Subsection(subsection).
		SetOption("headRef", request.HeadRef).
		SetOption("headRepo", request.HeadRepo).
		SetOption("baseRef", request.BaseRef).
		SetOption("baseRepo", request.BaseRepo).
		SetOption("number", strconv.Itoa(request.Number)).
		SetOption("webUrl", request.WebURL)

	err = c.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}
