

// Main Web Entrance
package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/cihub/seelog"
	"flag"
	"mime"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"rabbit/conf"
	"rabbit/controllers"
	"rabbit/lib"
	"rabbit/models"
	"rabbit/routers"
)

func init() {
	// init flag
	flags := conf.FlagConfig{}

	// user that hide
	flags.User = flag.String("user", "", "user")

	// db init or rebuild
	flags.DbInit = flag.Bool("db", false, "init db")
	flags.DbInitForce = flag.Bool("f", false, "force init db first drop db then rebuild it")

	// rbac config rebuild
	flags.Rbac = flag.Bool("rbac", false, "rebuild rbac database tables")

	// front-end  view
	home := flag.String("home", "", "home template")

	// config file position
	config := flag.String("config", "", "config file position if empty use default")

	flag.Parse()

	// init config
	if *config != "" {
		seelog.Trace("use diy config")
		err := beego.LoadAppConfig("ini", *config)
		if err != nil {
			seelog.Trace(err.Error())
		} else {
			seelog.Trace("Use config:" + *config)
		}
	}

	if *home != "" {
		seelog.Trace("Home template is " + *home)
		beego.AppConfig.Set(beego.BConfig.RunMode+"::"+"home_template", *home)
	}

	conf.InitConfig()

	// init lang
	// just add some ini in conf such locale_zh-CN.ini and edit app.conf
	langTypes := strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		seelog.Trace("Load language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			seelog.Error("Load language error:", err)
			return
		}
	}

	// add func map
	seelog.Trace("add i18n function map")
	beego.AddFuncMap("i18n", i18n.Tr)

	seelog.Trace("add stringsToJson function  map")
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	mime.AddExtensionType(".css", "text/css") // some not important

	// init model
	seelog.Trace("model run")
	models.Run(flags)

	// init router
	seelog.Trace("router run")
	routers.Run()

	seelog.Trace("start open error template")
	beego.ErrorController(&controllers.ErrorController{})
}

// Start!
func main() {
	seelog.Trace("kan")
	// 开启 ORM 调试模式
	orm.Debug = true 
	// 自动建表
	orm.RunSyncdb("default", false, true)
	seelog.Trace("Start Listen ...")
	beego.Run()
}
