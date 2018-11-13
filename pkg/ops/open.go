package ops

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/open-golang/open"
)

func Open(prProvider PRProvider) error {
	pr, err := prProvider.Get()
	if noPrErr, ok := err.(git.NoPRForBranch); ok {
		fmt.Printf("There is no pull request associated with branch '%s'\n", noPrErr.Branch)
		fmt.Println("    (use \"git pr fetch\" to get pull request if it exists)")
		fmt.Println("    (use \"git pr create\" to create pull request for current branch)")
		return err
	} else if err != nil {
		return err
	}

	err = open.Start(pr.WebURL)
	if err != nil {
		return err
	}

	return nil
}

type PRProvider interface {
	Get() (*git.PullRequest, error)
}

type PRByNumber struct {
	Repo   git.Repo
	Number int
}

func (p PRByNumber) Get() (*git.PullRequest, error) {
	prs, err := p.Repo.ListPRs()
	if err != nil {
		return nil, err
	}

	matches := make([]*git.PullRequest, 0)

	for _, pr := range prs {
		if pr.Number == p.Number {
			matches = append(matches, pr)
		}
	}

	if len(matches) == 0 {
		return nil, git.PRNotFound{Number: p.Number}
	}
	if len(matches) > 1 {
		remotes := make([]string, len(matches))
		for i, pr := range matches {
			remotes[i] = pr.Remote
		}
		return nil, errors.Errorf("ambiguous PR number found in remotes %+v", remotes)
	}

	return matches[0], nil
}

type PRByBranch struct {
	Repo   git.Repo
	Branch string
}

func (p PRByBranch) Get() (*git.PullRequest, error) {
	return p.Repo.GetPRForBranch(p.Branch)
}

type PRForCurrentBranch struct {
	Repo git.Repo
}

func (p PRForCurrentBranch) Get() (*git.PullRequest, error) {
	return getPRForCurrentBranch(p.Repo)
}
