package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/git"
	"strings"
)

func Status(repo git.Repo) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if _, ok := err.(git.NoPullRequestForBranch); err != nil && ok {
		fmt.Println("No pull request for current branch")
		return nil
	} else if err != nil {
		return err
	}

	fmt.Printf("On pull request %s#%v\n", pr.Remote, pr.Number)
	fmt.Printf("    %s\n", pr.Title)
	fmt.Println()
	if pr.Comment == "" {
		fmt.Printf("Pull request %s#%v is in sync with GitHub\n", pr.Remote, pr.Number)
	} else {
		fmt.Printf("Pull request %s#%v is not out-of-sync\n", pr.Remote, pr.Number)
		fmt.Printf(`    (use "git hub push" to push comment to GitHub)\n`)
		fmt.Println("Comment:")
		for _, line := range strings.Split(pr.Comment, "\n") {
			fmt.Printf("    %s\n", line)
		}
	}

	return nil
}
