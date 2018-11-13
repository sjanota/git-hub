package ops

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/github"
	"github.com/sjanota/open-golang/open"
)

func Create(repo git.Repo) error {
	branch, err := repo.GetCurrentBranch()
	if err != nil {
		return errors.Wrap(err, "cannot get current branch")
	}

	remote, err := repo.GetBranchRemote(branch)
	if err != nil {
		return errors.Wrapf(err, "cannot get remote for branch %s", branch)
	}

	remoteURL, err := repo.GetRemoteURL(remote)
	if err != nil {
		return err
	}

	url, err := github.ParseURL(remoteURL)
	if err != nil {
		return err
	}

	createUrl := fmt.Sprintf("https://github.com/%s/compare/%s?expand=1", url.Path, branch)

	return open.Start(createUrl)
}
