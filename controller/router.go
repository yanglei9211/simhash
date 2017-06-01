package controller

import (
	"github.com/astaxie/beego"
)

type UrlHandler struct {
	url			string
	controller	beego.ControllerInterface
}

var allUrls = []UrlHandler{
	UrlHandler{"/test", &TestGetter{}},
}