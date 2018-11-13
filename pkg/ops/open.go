package ops

import (
	"fmt"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/open-golang/open"
)

func Open(repo git.Repo) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if noPrErr, ok := err.(git.NoPullRequestForBranch); ok {
		fmt.Printf("There is no pull request associated with branch %s\n", noPrErr.Branch)
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
