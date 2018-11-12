package git_hub

import (
	"github.com/sjanota/git-hub/pkg/cli"
	"github.com/sjanota/git-hub/pkg/git"
)

func Commands(repo git.Repo) []cli.Command {
	return []cli.Command{
		&fetch{repo: repo},
		&clean{repo: repo},
		&open{repo: repo},
		&status{repo: repo},
		&comment{repo: repo},
		&push{repo: repo},
	}
}
