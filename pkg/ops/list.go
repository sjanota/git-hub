package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/git"
)

func List(repo git.Repo) error {
	prs, err := repo.ListPullRequests()
	if err != nil {
		return err
	}

	currentPr, err := getPullRequestForCurrentBranch(repo)
	if _, ok := err.(git.NoPullRequestForBranch); err != nil && !ok {
		return err
	}

	for _, pr := range prs {
		if currentPr != nil && pr.Number == currentPr.Number {
			fmt.Printf("* %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		} else {
			fmt.Printf("  %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		}
	}

	return nil
}
