package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/config"
)

func Status(cfg config.Config) error {
	prs, err := cfg.ListPullRequests()
	if err != nil {
		return err
	}

	currentPr, err := getPullRequestForCurrentBranch(cfg)
	if err != nil {
		return err
	}

	for _, pr := range prs {
		if pr.Number == currentPr.Number {
			fmt.Printf("*%6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		} else {
			fmt.Printf(" %6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		}
	}

	return nil
}
