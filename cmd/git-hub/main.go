package main

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/cli/git-hub"
	"github.com/sjanota/git-hub/pkg/git"
	"log"
	"os"
)

func main() {
	app := cli.App("git-hub", "Use GitHub from CLI")

	cfg, err := git.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, cmd := range git_hub.Commands(cfg) {
		cmd.Configure(app)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
