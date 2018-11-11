package github

import (
	"context"
	"github.com/google/go-github/v18/github"
)

type Client interface {
	GetPullRequests(owner, repo string, filter PullRequestFilter) (PullRequests, error)
}

type client struct {
	gh *github.Client
}

var _ Client = &client{}

func NewClient() Client {
	return &client{
		gh: github.NewClient(nil),
	}
}

func (c client) GetPullRequests(owner, repo string, filter PullRequestFilter) (PullRequests, error) {

	prs, _, err := c.gh.PullRequests.List(context.Background(), owner, repo, &github.PullRequestListOptions{
		State: "open",
	})

	if err != nil {
		return nil, err
	}

	return &pullRequests{prs: prs, filter: filter}, nil
}
