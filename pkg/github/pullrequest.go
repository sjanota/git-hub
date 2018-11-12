package github

import (
	"github.com/google/go-github/v18/github"
	"github.com/sjanota/git-pr/pkg/git"
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
	Iter() <-chan *git.PullRequest
}

type pullRequests struct {
	prs    []*github.PullRequest
	next   int
	filter PullRequestFilter
}

func convertPulRequest(pr *github.PullRequest) *git.PullRequest {
	return &git.PullRequest{
		HeadRef:  *pr.Head.Ref,
		HeadRepo: *pr.Head.Repo.FullName,
		Number:   *pr.Number,
		WebURL:   *pr.HTMLURL,
		Remote:   *pr.Base.Repo.FullName,
		Title:    *pr.Title,
		Comment:  "",
	}
}

func (prs *pullRequests) Iter() <-chan *git.PullRequest {
	ch := make(chan *git.PullRequest)
	go func() {
		for _, pr := range prs.prs {
			if prs.filter.filter(pr) {
				ch <- convertPulRequest(pr)
			}
		}
		close(ch)
	}()
	return ch
}
