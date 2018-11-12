package git_hub

import (
	"github.com/sjanota/git-hub/pkg/cli"
	"github.com/sjanota/git-hub/pkg/git"
)

func Commands(cfg git.Config) []cli.Command {
	return []cli.Command{
		&fetch{cfg: cfg},
		&clean{cfg: cfg},
		&open{cfg: cfg},
		&status{cfg: cfg},
		&comments{cfg: cfg},
	}
}
