package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/ops"
)

type Fetch struct {}

func (f *Fetch) Configure(app *cli.Cli) {
	app.Command("fetch", "Fetch Pull Requests", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := ops.FetchPullRequests()
			if err != nil {
				panic(err)
			}
		}
	})
}