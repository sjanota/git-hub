package git_hub

import (
	"github.com/sjanota/git-hub/pkg/cli"
	"github.com/sjanota/git-hub/pkg/config"
)

func Commands(cfg config.Config) []cli.Command {
	return []cli.Command{
		&fetch{cfg: cfg},
		&clean{},
		&open{cfg: cfg},
	}
}
