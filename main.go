package main // import "go.arkanosis.net/gogogo"

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
)

var VERSION = "unknown"

func main() {
	usage := `
Usage: gogogo <project>
       gogogo -h | --help
       gogogo --version

Arguments:
    project        Project name.

Options:
    -h, --help    Show this screen.
    --version     Show version.`

	opts, _ := docopt.ParseArgs(usage, os.Args[1:], "gogogo v"+VERSION)
	project, _ := opts.String("<project>")
	fmt.Println("Creating project `" + project + "'")
	os.Mkdir(project, 755)
}
