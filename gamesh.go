package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gamesh"
	app.Version = Version
	app.Usage = ""
	app.Author = "Daisuke Fujita"
	app.Email = "dtanshi45@gmail.com"
	app.Commands = Commands

	return app
}
