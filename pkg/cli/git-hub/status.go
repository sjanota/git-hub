package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/git"
	"github.com/sjanota/git-hub/pkg/ops"
)

type status struct {
	repo git.Repo
}

func (s *status) Configure(app *cli.Cli) {
	app.Command("status", "See status of pull requests", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := ops.Status(s.repo)
			if err != nil {
				panic(err)
			}
		}
	})
}
