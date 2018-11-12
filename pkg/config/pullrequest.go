package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4/plumbing/format/config"
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
}

const (
	pullRequestSection  = "github-pr"
	pullRequestTitle    = "title"
	pullRequestComment  = "comment"
	pullRequestHeadRef  = "headRef"
	pullRequestHeadRepo = "headRepo"
	pullRequestWebUrl   = "webUrl"
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

	cfg.Raw.Section(pullRequestSection).Subsection(subsection).
		SetOption(pullRequestTitle, request.Title).
		SetOption(pullRequestComment, request.Comment).
		SetOption(pullRequestHeadRef, request.HeadRef).
		SetOption(pullRequestHeadRepo, request.HeadRepo).
		SetOption(pullRequestWebUrl, request.WebURL)

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
		Title:    subsection.Option(pullRequestTitle),
		Comment:  subsection.Option(pullRequestComment),
		HeadRef:  subsection.Option(pullRequestHeadRef),
		HeadRepo: subsection.Option(pullRequestHeadRepo),
		Number:   number,
		WebURL:   subsection.Option(pullRequestWebUrl),
		Remote:   name[0],
	}, nil
}
