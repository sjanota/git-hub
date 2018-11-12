package git

import "fmt"

type PullRequestNotFound struct {
	Number int
	Remote string
}

var _ error = PullRequestNotFound{}

func (PullRequestNotFound) Error() string {
	return "pull request not found"
}

type NoPullRequestForBranch struct {
	Branch string
}

var _ error = NoPullRequestForBranch{}

func (e NoPullRequestForBranch) Error() string {
	return fmt.Sprintf("no pull request for branch '%s'", e.Branch)
}

type NoRemoteForURL struct {
	Url string
}

var _ error = NoRemoteForURL{}

func (e NoRemoteForURL) Error() string {
	return fmt.Sprintf("no remote for url '%s'", e.Url)
}
