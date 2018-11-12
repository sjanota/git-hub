package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/git"
)

type clean struct {
	cfg git.Repo
}

func (c *clean) Configure(app *cli.Cli) {
	app.Command("clean", "clean GitHub data from repo git", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := c.cfg.Clean()
			if err != nil {
				panic(err)
			}
		}
	})
}
