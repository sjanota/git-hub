package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/git-hub/pkg/github"
)

func Push(repo git.Repo) error {
	pr, err := getPullRequestForCurrentBranch(repo)
	if err != nil {
		return err
	}

	if pr.InSync {
		fmt.Println("Already up-to-date")
		return nil
	}

	remoteUrl, err := github.RepoURL(pr.Remote)
	if err != nil {
		return err
	}

	remote, err := repo.GetRemoteForURL(remoteUrl.Full)
	if err != nil {
		return err
	}

	credentials, err := repo.GetCredentials(remote)
	if err != nil {
		return err
	}

	gh := github.NewClient(credentials)
	err = gh.PushPullRequestComment(remoteUrl, pr)
	if err != nil {
		return err
	}

	pr.InSync = true
	return repo.StorePullRequest(pr)
}
