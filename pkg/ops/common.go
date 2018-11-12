package ops

import "github.com/sjanota/git-hub/pkg/git"

func getPullRequestForCurrentBranch(cfg git.Repo) (*git.PullRequest, error) {
	branch, err := cfg.GetCurrentBranch()
	if err != nil {
		return nil, err
	}

	return cfg.GetPullRequestForBranch(branch)
}
