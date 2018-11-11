package ops

import (
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/git-hub/pkg/github"
)

func Fetch() {

}

func FetchPullRequests() error {
	gh := github.NewClient()
	cfg, err := config.NewGitConfig()
	if err != nil {
		return err
	}

	prs, err := gh.GetPullRequests("kyma-project", "kyma", github.PullRequestFilter{
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

	return nil
}
