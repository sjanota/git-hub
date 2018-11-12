package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/git-hub/pkg/ops"
)

type fetch struct {
	cfg    config.Config
	remote *string
	all    *bool
}

func (f *fetch) Configure(app *cli.Cli) {
	app.Command("fetch f", "fetch Pull Requests", func(cmd *cli.Cmd) {
		cmd.Spec = "[REMOTE | -a]"
		f.remote = cmd.StringArg("REMOTE", "origin", "Optional remote name to fetch")
		f.all = cmd.BoolOpt("all a", false, "fetch all remotes")

		cmd.Action = f.action
	})
}

func (f *fetch) action() {
	var remotes config.RemotesLister
	if *f.all {
		remotes = config.AllRemotesLister{}
	} else {
		remotes = config.OneRemoteLister{Remote: *f.remote}
	}

	err := ops.FetchPullRequests(f.cfg, remotes)
	if err != nil {
		panic(err)
	}
}
