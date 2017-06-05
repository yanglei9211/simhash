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
	ss := []string{"放大是否打算"}
	res := []simhash.Simhash{}
	for _, s := range ss {
		sim := simhash.Simhash{}
		sim.Init(s)
		res = append(res, sim)
	}
	for _, r := range res {
		fmt.Println(r.Value())
	}
	n0 := simhash.IndexNode{}
	n0.Init(res[0], "0")
	test := []simhash.IndexNode{n0}
	s := simhash.SimhashIndex{}
	s.Init(test)
}
