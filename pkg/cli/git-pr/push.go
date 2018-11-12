package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type push struct {
	repo git.Repo
}

func (p *push) Configure(app *cli.Cli) {
	app.Command("push", "Push comment to GitHub", func(cmd *cli.Cmd) {
		cmd.Action = p.action
	})
}

func (p push) action() {
	err := ops.Push(p.repo)
	if err != nil {
		panic(err)
	}
}
