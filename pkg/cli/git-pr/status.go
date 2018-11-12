package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type status struct {
	repo git.Repo
}

func (s *status) Configure(app *cli.Cli) {
	app.Command("status", "See status of pull requests", func(cmd *cli.Cmd) {
		cmd.Action = s.action
	})
}

func (s status) action() {
	err := ops.Status(s.repo)
	if err != nil {
		panic(err)
	}
}
