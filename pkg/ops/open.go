package ops

import (
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/open-golang/open"
)

func Open(cfg git.Repo) error {
	pr, err := getPullRequestForCurrentBranch(cfg)
	if err != nil {
		return err
	}

	err = open.Start(pr.WebURL)
	if err != nil {
		return err
	}

	return nil
}
