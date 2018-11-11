package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
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
		cmd.Spec = "[-a] [REMOTE] [-a]"
		f.remote = cmd.StringArg("REMOTE", "origin", "Optional remote name to fetch")
		f.all = cmd.BoolOpt("all a", false, "fetch all remotes")

		cmd.Action = f.action
	})
}

func (f *fetch) action() {
	var remotes []string
	var err error
	if *f.all {
		remotes, err = f.cfg.ListRemoteNames()
		if err != nil {
			panic(errors.Wrap(err, "cannot list remotes"))
		}
	} else {
		remotes = []string{*f.remote}
	}

	for _, remote := range remotes {
		err := ops.FetchPullRequests(f.cfg, remote)
		if err != nil {
			panic(err)
		}
	}
}
