package flag

import (
	"flag"
	"os"
	"testing"
)

var name string

func TestFlag(t *testing.T) {
	flag.Parse()

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("phh", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	args := os.Args
	switch args[0] {
	case "go":

	}
}
