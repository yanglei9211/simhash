package controller

import (
	"fmt"
)

type TestGetter struct {
	BasicController
}

func (sg *TestGetter) Get() {
	fmt.Println("get!!!!")
	res := "aaabbbccc"
	sg.writeReponse(map[string]interface{}{
		"item": res,
	})
}

func (sg * TestGetter) Post() {
	fmt.Println("post!!!!!")
	res := "cccbbbaaa"
	sg.writeReponse(map[string]interface{}{
		"item": res,
	})
}
