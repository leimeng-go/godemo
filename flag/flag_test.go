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
		t.Logf("")
	}
}

// go test -v -args "value hello" -run ^TestHello$ flag_test.go  

func TestHello(t *testing.T) {
	append(os.Args, )
	t.Logf("参数内容: %+v", os.Args)
	var name string
	flag.StringVar(&name,"name", "333", "这是一个测试命令")
	flag.Parse()
	t.Log(name)
}
