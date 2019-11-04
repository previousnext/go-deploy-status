package main

import (
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/previousnext/go-deploy-status/cmd/create"
)

func main() {
	app   := kingpin.New("deploy-status", "A command-line tool for interacting with GitHub deployment API")
	createCmd := app.Command("create", "Create a new GitHub API object. See sub-commands for options.")
	create.Deployment(createCmd)
	create.Status(createCmd)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
