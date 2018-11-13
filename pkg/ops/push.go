package ops

import (
	"fmt"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/github"
)

func Push(repo git.Repo) error {
	pr, err := getPRForCurrentBranch(repo)
	if err != nil {
		return err
	}

	if pr.Comment == "" {
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

	pr.Comment = ""
	return repo.StorePR(pr)
}
