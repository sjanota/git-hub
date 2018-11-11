package ops

import (
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/open-golang/open"
)

func Open(cfg config.Config) error {
	branch, err := cfg.GetCurrentBranch()
	if err != nil {
		return err
	}

	pr, err := cfg.GetPullRequestForBranch(branch)
	if err != nil {
		return err
	}

	err = open.Start(pr.WebURL)
	if err != nil {
		return err
	}

	return nil
}
