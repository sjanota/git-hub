package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/ops"
)

type Clean struct{}

func (c *Clean) Configure(app *cli.Cli) {
	app.Command("clean", "Clean GitHub data from repo config", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := ops.Clean()
			if err != nil {
				panic(err)
			}
		}
	})
}
