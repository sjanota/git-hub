package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/git-hub/pkg/ops"
)

type status struct {
	cfg config.Config
}

func (s *status) Configure(app *cli.Cli) {
	app.Command("status", "See status of pull requests", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			err := ops.Status(s.cfg)
			if err != nil {
				panic(err)
			}
		}
	})
}
