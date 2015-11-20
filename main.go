package main

import (
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gitlab.com/gitlab-org/gitlab-ci-multi-runner/helpers/cli"

	"fmt"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/commands"
	"gitlab.com/gitlab-org/gitlab-ci-multi-runner/common"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/executors/docker"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/executors/parallels"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/executors/shell"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/executors/ssh"
	_ "gitlab.com/gitlab-org/gitlab-ci-multi-runner/shells"
)

var NAME = "gitlab-ci-multi-runner"
var VERSION = "dev"
var REVISION = "HEAD"

func init() {
	common.NAME = NAME
	common.VERSION = VERSION
	common.REVISION = REVISION
}

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "a GitLab Runner"
	app.Version = fmt.Sprintf("%s (%s)", common.VERSION, common.REVISION)
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamil Trzciński",
			Email: "ayufan@ayufan.eu",
		},
	}
	cli_helpers.SetupLogLevelOptions(app)
	app.Commands = common.GetCommands()
	app.CommandNotFound = func(context *cli.Context, command string) {
		log.Fatalln("Command", command, "not found.")
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
