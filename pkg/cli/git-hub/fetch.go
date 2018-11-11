package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/ops"
)

type Fetch struct {
	remote *string
}

func (f *Fetch) Configure(app *cli.Cli) {
	app.Command("fetch", "Fetch Pull Requests", func(cmd *cli.Cmd) {
		cmd.Spec = "[REMOTE]"
		f.remote = cmd.StringArg("REMOTE", "origin", "Optional remote name to fetch")
		cmd.Action = func() {
			err := ops.FetchPullRequests(*f.remote)
			if err != nil {
				panic(err)
			}
		}
	})
}
