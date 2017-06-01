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
	s1 := "发觉肯定是离开家理发店addfadf沙发的设计科罗拉多放假开始54564654654"
	ts := simhash.Simhash{}
	ts.Init(s1)
	res := ts.Slide()
	fmt.Println(res)
}