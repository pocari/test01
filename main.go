package main

import (
	"fmt"
	"github.com/pocari/gosample"
	"github.com/pocari/test01/util"
)

func main() {
	fmt.Println("hello, world")
	fmt.Println(gosample.Message)
	fmt.Println(gosample.Message3)
	fmt.Printf("hoge is %v\n", util.TestFunc(1, 2))
}
