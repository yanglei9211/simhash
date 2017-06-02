package main

import (
	/*
	"cluster/controller"
	"runtime"
	"github.com/astaxie/beego"
	*/
	"cluster/simhash"
	"fmt"
)

func main() {
	/*
	controller.Init()
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
	*/
	s1 := "54564654654"
	ts := simhash.Simhash{}
	ts.Init(s1)
	fmt.Println(ts.Value())
}