package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGet,
	commandList,
}

var commandGet = cli.Command{
	Name:  "get",
	Usage: "",
	Description: `
`,
	Action: doGet,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: `
`,
	Action: doList,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doGet(c *cli.Context) {
}

func doList(c *cli.Context) {
}
