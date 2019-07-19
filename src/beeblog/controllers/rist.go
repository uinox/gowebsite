package controllers

import (
	"beeblog/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type RistController struct {
	beego.Controller
}

func (user *RistController) Rist() {
	ob := &models.Tm{}
	json.Unmarshal(user.Ctx.Input.RequestBody, ob)
	fmt.Println("-----------")
	fmt.Println(*ob)
	fmt.Println("-----------")
	if ob.Name == "any" {
		// user.Data["json"] = map[string]interface{}{"interest": "football", "major": "computer"}
	}
	if ob.Name == "mark" {
		// user.Data["json"] = map[string]interface{}{"interest": "travel", "major": "economics"}
	}

	user.ServeJSON()

	return

}
