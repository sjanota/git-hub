package main

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-hub/pkg/cli/git-hub"
	"log"
	"os"
)

func main() {
	app := cli.App("git-hub", "Use GitHub from CLI")

	for _, cmd := range git_hub.Commands() {
		cmd.Configure(app)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
