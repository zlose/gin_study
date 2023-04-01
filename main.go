package main

import (
	"fmt"
	"gin_study/core"
	"gin_study/global"
)

func main() {
	core.Initconf()
	fmt.Println(global.Config)
}
