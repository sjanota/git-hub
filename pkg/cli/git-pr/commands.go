package git_pr

import (
	"github.com/sjanota/git-pr/pkg/cli"
	"github.com/sjanota/git-pr/pkg/git"
)

func Commands(repo git.Repo) []cli.Command {
	return []cli.Command{
		&fetch{repo: repo},
		&clean{repo: repo},
		&open{repo: repo},
		&status{repo: repo},
		&comment{repo: repo},
		&push{repo: repo},
		&root{repo: repo},
		&create{repo: repo},
	}
}
