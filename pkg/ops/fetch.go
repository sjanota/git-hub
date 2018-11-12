package ops

import (
	"github.com/pkg/errors"
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/git-hub/pkg/github"
)

func FetchPullRequests(cfg config.Config, remotesLister config.RemotesLister) error {
	gh := github.NewClient()
	remotes, err := remotesLister.List(cfg)
	if err != nil {
		return err
	}

	for _, remote := range remotes {

		remoteUrl, err := cfg.GetRemoteURL(remote)
		if err != nil {
			return errors.Wrapf(err, "cannot get remote url %s", remote)
		}

		url, err := github.ParseURL(remoteUrl)
		if err != nil {
			return errors.Wrapf(err, "cannot parse remote url %s", remoteUrl)
		}

		prs, err := gh.GetPullRequests(url.Owner, url.RepositoryName, github.PullRequestFilter{
			AssigneeLogin: "sjanota",
		})

		if err != nil {
			return err
		}

		for pr := range prs.Iter() {
			err := cfg.StorePullRequest("kyma", pr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
