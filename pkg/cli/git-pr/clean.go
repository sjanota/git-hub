package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
)

type clean struct {
	repo git.Repo
}

func (c *clean) Configure(app *cli.Cli) {
	app.Command("clean", "clean GitHub data from repo git", func(cmd *cli.Cmd) {
		cmd.Action = c.action
	})
}

func (c clean) action() {
	err := c.repo.Clean()
	if err != nil {
		panic(err)
	}
}
