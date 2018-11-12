package ops

import (
	"github.com/sjanota/git-hub/pkg/git"
)

func Comment(repo git.Repo, editor git.CommentEditor) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if err != nil {
		return err
	}

	comment, err := editor.Edit(pr)
	if err != nil {
		return err
	}

	pr.Comment = comment

	return repo.StorePullRequest(pr)
}
