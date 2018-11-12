package ops

import "github.com/sjanota/git-hub/pkg/git"

func Comment(cfg git.Config, comment string) error {
	pr, err := getPullRequestForCurrentBranch(cfg)
	if err != nil {
		return err
	}

	pr.Comment = comment

	return cfg.StorePullRequest(pr)
}
