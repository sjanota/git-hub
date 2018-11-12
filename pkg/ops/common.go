package ops

import "github.com/sjanota/git-hub/pkg/git"

func getPullRequestForCurrentBranch(repo git.Repo) (*git.PullRequest, error) {
	branch, err := repo.GetCurrentBranch()
	if err != nil {
		return nil, err
	}

	return repo.GetPullRequestForBranch(branch)
}
