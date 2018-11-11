package git_hub

import "github.com/sjanota/git-hub/pkg/cli"

func Commands() []cli.Command {
	return []cli.Command{
		&Fetch{},
		&Clean{},
	}
}
