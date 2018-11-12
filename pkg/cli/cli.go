package cli

import "github.com/jawher/mow.cli"

type Command interface {
	Configure(app *cli.Cli)
}
