package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/git-hub/pkg/ops"
)

type open struct {
	cfg git.Repo
}

func (o *open) Configure(app *cli.Cli) {
	app.Command("open", "Opens pull request page in your default browser", func(cmd *cli.Cmd) {
		cmd.Action = o.action
	})
}

func (o *open) action() {
	err := ops.Open(o.cfg)
	if err != nil {
		panic(err)
	}
}
