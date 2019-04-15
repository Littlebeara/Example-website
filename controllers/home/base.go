
package home

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"rabbit/conf"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	this.Lang = ""

	l := this.Ctx.GetCookie("lang")
	if l == "" {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if i18n.IsExist(al) {
				this.Lang = al
			}
		}
	} else {
		switch l {
		case "en":
			this.Lang = "en-US"
		default:
			this.Lang =  "zh-CN"

		}
	}
	if this.Lang == "" {
		this.Lang = "zh-CN"
	}

	this.Data["Lang"] = this.Lang

	// Add some mark
	if this.Ctx.GetCookie("X-Home") == "" {
		this.Ctx.SetCookie("X-Home", conf.HomeTemplate, false, "/", false, false, true)
	}

	this.Ctx.Output.Header("X-Version", conf.Version)
}

// Get the address of template
func (this *baseController) GetTemplate() string {
	templatetype := conf.HomeTemplate

	temp := this.Ctx.GetCookie("X-Home")
	if temp != "" {
		// todo dangerous
		// magic way for me to change
		switch temp {
		case "home/first", "home/default", "home/hunterhug":
			templatetype = temp
		default:
			break
		}

	}
	return templatetype
}

// json response
// if false will be 404
func (this *baseController) Rsp(status bool, str string) {
	if status {
		this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
		this.ServeJSON()
	}
	this.StopRun()
}
