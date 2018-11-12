package github

import (
	"context"
	"github.com/google/go-github/v18/github"
	"github.com/sjanota/git-pr/pkg/git"
	"net/http"
	"time"
)

type Client interface {
	GetPullRequests(url *URL, filter PullRequestFilter) (PullRequests, error)
	PushPullRequestComment(url *URL, pr *git.PullRequest) error
}

type client struct {
	gh *github.Client
}

var _ Client = &client{}

func NewClient(credentials *git.Credentials) Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &github.BasicAuthTransport{
			Username: credentials.Username,
			Password: credentials.Password,
		},
	}

	return &client{
		gh: github.NewClient(httpClient),
	}
}

func (c client) PushPullRequestComment(url *URL, pr *git.PullRequest) error {
	_, _, err := c.gh.Issues.CreateComment(context.Background(), url.Owner, url.RepositoryName, pr.Number, &github.IssueComment{
		Body: &pr.Comment,
	})
	return err
}

func (c client) GetPullRequests(url *URL, filter PullRequestFilter) (PullRequests, error) {

	prs, _, err := c.gh.PullRequests.List(context.Background(), url.Owner, url.RepositoryName, &github.PullRequestListOptions{
		State: "open",
	})

	if err != nil {
		return nil, err
	}

	return &pullRequests{prs: prs, filter: filter}, nil
}
