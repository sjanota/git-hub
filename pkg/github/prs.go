package github

import "github.com/google/go-github/v18/github"

type PullRequestFilter struct {
	AssigneeLogin string
}

func (f PullRequestFilter) Filter(pr *github.PullRequest) bool {
	if f.assigneeLoginMismatch(pr) {
		return false
	}

	return true
}

func (f PullRequestFilter) assigneeLoginMismatch(pr *github.PullRequest) bool {
	return f.AssigneeLogin != "" && (pr.Assignee == nil || *pr.Assignee.Login != f.AssigneeLogin)
}

type PullRequests interface {
	Iter() <-chan *github.PullRequest
}

type pullRequests struct {
	prs    []*github.PullRequest
	next   int
	filter PullRequestFilter
}

func (prs *pullRequests) Iter() <-chan *github.PullRequest {
	ch := make(chan *github.PullRequest)
	go func() {
		for _, pr := range prs.prs {
			if prs.filter.Filter(pr) {
				ch <- pr
			}
		}
		close(ch)
	}()
	return ch
}
