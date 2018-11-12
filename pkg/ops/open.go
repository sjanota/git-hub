package ops

import (
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/open-golang/open"
)

func Open(cfg config.Config) error {
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
