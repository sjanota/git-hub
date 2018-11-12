package ops

import (
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/open-golang/open"
)

func Open(repo git.Repo) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if err != nil {
		return err
	}

	err = open.Start(pr.WebURL)
	if err != nil {
		return err
	}

	return nil
}
