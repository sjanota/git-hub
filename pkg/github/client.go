package github

import "github.com/google/go-github/v18/github"

type Client interface {
	GetMyPullRequests() ([]*github.PullRequest, error)
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

func (c client) GetMyPullRequests() ([]*github.PullRequest, error) {

	prs, _, err := c.gh.PullRequests.List(nil, "kyma-project", "kyma", &github.PullRequestListOptions{
		State: "open",
		Head:  "assignee:sjanota",
	})

	if err != nil {
		return nil, err
	}

	return prs, nil
}