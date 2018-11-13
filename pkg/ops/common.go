package ops

import "github.com/sjanota/git-pr/pkg/git"

func getPRForCurrentBranch(repo git.Repo) (*git.PullRequest, error) {
	branch, err := repo.GetCurrentBranch()
	if err != nil {
		return nil, err
	}

	return repo.GetPRForBranch(branch)
}
