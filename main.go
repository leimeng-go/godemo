package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main(){
	
	var name string
	flag.StringVar(&name,"name","默认值","简单测试命令行参数解析")
	flag.Parse()
	log.Printf("命令行输入参数: %+v",os.Args)
	fmt.Println(name)
}