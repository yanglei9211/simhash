package controller

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type BasicController struct {
	beego.Controller
}

func (b *BasicController) writeReponse(r map[string]interface{}) {
	response, err := json.Marshal(
		map[string]interface{}{
			"status": 1,
			"data": r,
		})
	if err != nil {
		panic(err)
	}
	b.Ctx.WriteString(string(response))
}