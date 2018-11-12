package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/git-hub/pkg/ops"
)

type root struct {
	repo git.Repo
}

func (r *root) Configure(app *cli.Cli) {
	app.Action = r.action
}

func (r root) action() {
	err := ops.List(r.repo)
	if err != nil {
		panic(err)
	}
}
