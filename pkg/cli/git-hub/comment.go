package git_hub

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/config"
	"github.com/sjanota/git-hub/pkg/ops"
)

type comments struct {
	comment *string
	cfg     config.Config
}

func (c *comments) Configure(app *cli.Cli) {
	app.Command("comment cm", "Edit pul request comment", func(cmd *cli.Cmd) {
		cmd.Spec = "-m"
		c.comment = cmd.StringOpt("comment m", "", "Text of the comment")
		cmd.Action = c.action
	})
}

func (c comments) action() {
	err := ops.Comment(c.cfg, *c.comment)
	if err != nil {
		panic(err)
	}
}
