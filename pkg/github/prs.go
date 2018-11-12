package github

import (
	"github.com/google/go-github/v18/github"
	"github.com/sjanota/git-hub/pkg/config"
)

type PullRequestFilter struct {
	AssigneeLogin string
}

func (f PullRequestFilter) filter(pr *github.PullRequest) bool {
	if f.assigneeLoginMismatch(pr) {
		return false
	}

	return true
}

func (f PullRequestFilter) assigneeLoginMismatch(pr *github.PullRequest) bool {
	return f.AssigneeLogin != "" && (pr.Assignee == nil || *pr.Assignee.Login != f.AssigneeLogin)
}

type PullRequests interface {
	Iter() <-chan *config.PullRequest
}

type pullRequests struct {
	prs    []*github.PullRequest
	next   int
	filter PullRequestFilter
}

func (prs *pullRequests) Iter() <-chan *config.PullRequest {
	ch := make(chan *config.PullRequest)
	go func() {
		for _, pr := range prs.prs {
			if prs.filter.filter(pr) {
				prConfig := &config.PullRequest{
					HeadRef:  *pr.Head.Ref,
					HeadRepo: *pr.Head.Repo.FullName,
					Number:   *pr.Number,
					WebURL:   *pr.HTMLURL,
					Remote:   *pr.Base.Repo.FullName,
					Title:    *pr.Title,
				}
				ch <- prConfig
			}
		}
		close(ch)
	}()
	return ch
}
