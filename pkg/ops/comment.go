package ops

import "github.com/sjanota/git-hub/pkg/config"

func Comment(cfg config.Config, comment string) error {
	pr, err := getPullRequestForCurrentBranch(cfg)
	if err != nil {
		return err
	}

	pr.Comment = comment

	return cfg.StorePullRequest(pr)
}
