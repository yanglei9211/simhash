package main

import (
	"simhash/controller"
	"runtime"
	"github.com/astaxie/beego"
)

func main() {
	controller.Init()
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}