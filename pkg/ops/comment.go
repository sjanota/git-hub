package ops

import (
	"fmt"
	"github.com/sjanota/git-pr/pkg/git"
)

func Comment(repo git.Repo, editor git.CommentEditor) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if _, ok := err.(git.NoPullRequestForBranch); ok {
		fmt.Println("No pull request for current branch")
		return nil
	} else if err != nil {
		return err
	}

	comment, err := editor.Edit(pr)
	if err != nil {
		return err
	}

	pr.Comment = comment

	return repo.StorePullRequest(pr)
}
