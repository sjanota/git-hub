package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type create struct{
	repo git.Repo
}

func (c *create) Configure(app *cli.Cli) {
	app.Command("create cr", "Opens page with pul request creation", func(cmd *cli.Cmd) {
		cmd.Action = c.action
	})
}

func (c create) action() {
	err := ops.Create(c.repo)
	if err != nil {
		panic(err)
	}
}


