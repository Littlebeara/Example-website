
package conf

import (
	//"fmt"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/cihub/seelog"
	"github.com/hunterhug/parrot/util"
)

type FlagConfig struct {
	User        *string
	DbInit      *bool
	DbInitForce *bool
	Rbac        *bool
}

var (
	AuthType      int
	AuthGateWay   string
	AuthAdmin     string
	Cookie7       bool
	Version       string
	HomeTemplate  string
	AdminTemplate string

	DbType   string
	DbHost   string
	DbName   string
	DbLog    string
	
	ConfigDir = util.CurDir()
)

func InitConfig() {
	// version
	Version = beego.AppConfig.DefaultString("version", "version2.0")
	seelog.Trace("Version:",Version)

	AuthType, _ = strconv.Atoi(beego.AppConfig.String("user_auth_type"))
	AuthGateWay = beego.AppConfig.DefaultString("rbac_auth_gateway", "/public/login")
	Cookie7, _ = beego.AppConfig.Bool("cookie7")
	AuthAdmin = beego.AppConfig.DefaultString("rbac_admin_user", "admin")
	HomeTemplate = beego.AppConfig.DefaultString("home_template", "default")

	DbType = beego.AppConfig.String("db_type")
	DbHost = beego.AppConfig.String("db_host")
	DbName = beego.AppConfig.String("db_name")
	DbLog = beego.AppConfig.String("dblog")

	

	AdminTemplate = beego.AppConfig.DefaultString("admin_template", "default")
}

func ForTestInitConfig() {
	err := beego.LoadAppConfig("ini", filepath.Join(ConfigDir, "app.conf"))
	if err != nil {
		panic(err.Error())
	}
	InitConfig()
}
