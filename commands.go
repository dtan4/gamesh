package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
	resp, err := http.Get("http://tokyo-ame.jwa.or.jp/scripts/mesh_index.js")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	indexList := strings.Replace(string(body), "Amesh.setIndexList(", "", -1)
	indexList = strings.Replace(indexList, ");", "", -1)
	indexList = strings.Replace(indexList, "\n", "", -1)
	indexList = strings.Replace(indexList, "\"", "", -1)
	indexList = strings.Replace(indexList, "[", "", -1)
	indexList = strings.Replace(indexList, "]", "", -1)
	indexList = strings.Replace(indexList, ",", "\n", -1)

	fmt.Println(indexList)
}
