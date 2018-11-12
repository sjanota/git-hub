package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/config"
	"strings"
)

const (
	statusCommentHeaderPadding = "└──"
	statusCommentPadding       = "   "
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
			fmt.Printf("* %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		} else {
			fmt.Printf("  %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		}

		if pr.Comment != "" {
			lines := strings.Split(pr.Comment, "\n")
			fmt.Println(statusCommentHeaderPadding, lines[0])
			for _, line := range lines[1:] {
				fmt.Println(statusCommentPadding, line)
			}
		}
	}

	return nil
}
