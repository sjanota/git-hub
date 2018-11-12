package ops

import "github.com/sjanota/git-hub/pkg/config"

func getPullRequestForCurrentBranch(cfg config.Config) (*config.PullRequest, error) {
	branch, err := cfg.GetCurrentBranch()
	if err != nil {
		return nil, err
	}

	return cfg.GetPullRequestForBranch(branch)
}
