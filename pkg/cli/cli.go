package cli

import "github.com/jawher/mow.cli"

type Command interface {
	Configure(cli * cli.Cli)
}
