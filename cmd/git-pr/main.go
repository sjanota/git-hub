package main

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/cli/git-pr"
	"github.com/sjanota/git-pr/pkg/git"
	"log"
	"os"
)

func main() {
	app := cli.App("git-pr", "Use GitHub pull requests from command line")

	cfg, err := git.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, cmd := range git_pr.Commands(cfg) {
		cmd.Configure(app)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
