package git

import "fmt"

type PRNotFound struct {
	Number int
	Remote string
}

var _ error = PRNotFound{}

func (PRNotFound) Error() string {
	return "pull request not found"
}

type NoPRForBranch struct {
	Branch string
}

var _ error = NoPRForBranch{}

func (e NoPRForBranch) Error() string {
	return fmt.Sprintf("no pull request for branch '%s'", e.Branch)
}

type NoRemoteForURL struct {
	Url string
}

var _ error = NoRemoteForURL{}

func (e NoRemoteForURL) Error() string {
	return fmt.Sprintf("no remote for url '%s'", e.Url)
}
