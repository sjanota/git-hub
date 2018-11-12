package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type open struct {
	repo git.Repo
}

func (o *open) Configure(app *cli.Cli) {
	app.Command("open", "Opens pull request page in your default browser", func(cmd *cli.Cmd) {
		cmd.Action = o.action
	})
}

func (o *open) action() {
	err := ops.Open(o.repo)
	if err != nil {
		panic(err)
	}
}
