package file

import (
	"fmt"
	"os"
)

func GetFileContent() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	file.Name()
}
func ls() {
	os.Open()
}
