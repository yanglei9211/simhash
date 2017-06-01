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

func bitsCount(num uint64) int {
	num = num - ((num >> 1) & 0x5555555555555555)
	num = (num & 0x3333333333333333) + ((num >> 2) & 0x3333333333333333)
	return int((((num + (num >> 4)) & 0xF0F0F0F0F0F0F0F) * 0x101010101010101) >> 56)
}