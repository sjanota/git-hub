package ops

import (
	"fmt"
	"github.com/sjanota/git-pr/pkg/git"
	"strings"
)

const (
	listCommentHeaderPadding = "  └──"
	listCommentPadding       = "     "
)

func List(repo git.Repo, details bool) error {
	prs, err := repo.ListPRs()
	if err != nil {
		return err
	}

	currentPr, err := getPRForCurrentBranch(repo)
	if _, ok := err.(git.NoPRForBranch); err != nil && !ok {
		return err
	}

	for _, pr := range prs {
		if currentPr != nil && pr.Number == currentPr.Number {
			fmt.Printf("* %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		} else {
			fmt.Printf("  %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		}

		if details && pr.Comment != "" {
			lines := strings.Split(pr.Comment, "\n")
			fmt.Println(listCommentHeaderPadding, lines[0])
			for _, line := range lines[1:] {
				fmt.Println(listCommentPadding, line)
			}
		}
	}

	return nil
}
