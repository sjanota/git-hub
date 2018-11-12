package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"strconv"
	"strings"
)

func (c *gitConfig) ListPullRequests() ([]*PullRequest, error) {
	cfg, err := c.repo.Config()
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

func (c *gitConfig) StorePullRequest(remote string, request *PullRequest) error {
	cfg, err := c.repo.Config()
	if err != nil {
		return err
	}

	subsection := fmt.Sprintf("%s:%v", remote, request.Number)

	cfg.Raw.Section("github-pr").Subsection(subsection).
		SetOption("title", request.Title).
		SetOption("headRef", request.HeadRef).
		SetOption("headRepo", request.HeadRepo).
		SetOption("number", strconv.Itoa(request.Number)).
		SetOption("webUrl", request.WebURL)

	err = c.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
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
		Title:    subsection.Option("title"),
		HeadRef:  subsection.Option("headRef"),
		HeadRepo: subsection.Option("headRepo"),
		Number:   number,
		WebURL:   subsection.Option("webUrl"),
		Remote:   name[0],
	}, nil
}
