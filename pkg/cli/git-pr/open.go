package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type open struct {
	repo git.Repo

	pr     *int
	branch *string
}

func (o *open) Configure(app *cli.Cli) {
	app.Command("open", "Opens pull request page in your default browser", func(cmd *cli.Cmd) {
		cmd.Spec = "[-p | -b]"

		o.pr = cmd.IntOpt("pr p", 0, "Number of PR to open")
		o.branch = cmd.StringOpt("branch b", "", "Number of PR to open")

		cmd.Action = o.action
	})
}

func (o *open) action() {
	var prProvider ops.PullRequestProvider
	if *o.pr != 0 {
		prProvider = ops.NumberPullRequestProvider{Repo: o.repo, Number: *o.pr}
	} else if *o.branch != "" {
		prProvider = ops.BranchPullRequestProvider{Repo: o.repo, Branch: *o.branch}
	} else {
		prProvider = ops.CurrentBranchPullRequestProvider{Repo: o.repo}
	}

	err := ops.Open(prProvider)
	if err != nil {
		panic(err)
	}
}
