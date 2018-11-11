package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/ops"
)

type clean struct{}

func (c *clean) Configure(app *cli.Cli) {
	app.Command("clean", "clean GitHub data from repo config", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := ops.Clean()
			if err != nil {
				panic(err)
			}
		}
	})
}
