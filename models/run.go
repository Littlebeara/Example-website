
package models

import (
	"github.com/cihub/seelog"
	"flag"
	"os"
	//"strings"

	//"github.com/hunterhug/marmot/miner"
	"rabbit/conf"
	"rabbit/models/util"
	"rabbit/models/admin"
	
)

func Run(config conf.FlagConfig) {
	seelog.Trace("database start to run")
	initDb(config)
	util.Connect()
	admin.InitData()
	// preRun(config) // I think maybe rid it off
}

func initDb(config conf.FlagConfig) {
	
	if !flag.Parsed() {
		flag.Parse()
	}
	if *config.DbInit {
		util.Syncdb(*config.DbInitForce)
		os.Exit(0)
	}
	if *config.Rbac {
		util.UpdateRbac()
		os.Exit(0)
	}
}


