package ops

import (
	"github.com/pkg/errors"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/github"
)

func FetchPullRequests(repo git.Repo, remotesLister git.RemotesLister) error {
	remotes, err := remotesLister.List(repo)
	if err != nil {
		return err
	}

	for _, remote := range remotes {
		credentials, err := repo.GetCredentials(remote)
		if err != nil {
			return err
		}

		gh := github.NewClient(credentials)

		remoteUrl, err := repo.GetRemoteURL(remote)
		if err != nil {
			return errors.Wrapf(err, "cannot get remote url %s", remote)
		}

		url, err := github.ParseURL(remoteUrl)
		if err != nil {
			return errors.Wrapf(err, "cannot parse remote url %s", remoteUrl)
		}

		prs, err := gh.GetPullRequests(url, github.PullRequestFilter{
			AssigneeLogin: credentials.Username,
		})
		if err != nil {
			return errors.Wrap(err, "cannot get pull requests")
		}

		for pr := range prs.Iter() {
			oldPr, err := repo.GetPullRequest(url.Path, pr.Number)
			if _, ok := err.(git.PullRequestNotFound); !ok {
				return err
			}
			if oldPr != nil {
				pr.Comment = oldPr.Comment
			}

			err = repo.StorePullRequest(pr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
