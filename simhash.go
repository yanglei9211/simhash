package main

import (
	"runtime"
	"github.com/astaxie/beego"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}