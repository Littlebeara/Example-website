package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"rabbit/models/admin"
	"rabbit/conf"
)
const (
	
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
}

func Createtb() {
	admin.InitData()
}

func Syncdb(force bool) {
	Createdb(force)
	Connect()
	Createtb()
	admin.InitData()
}

func UpdateRbac() {
	TrunCateRbacTable([]string{beego.AppConfig.String("rbac_group_table"), beego.AppConfig.String("rbac_node_table")})
	Connect()
	admin.InsertGroup()
	admin.InsertNodes()
}

func Createdb(force bool) {
	var createdbsql, dropdbsql string

	switch conf.DbType {
	case "sqlite3":
		createdbsql = fmt.Sprintf("CREATE　DATABASE　if not exits `%s` CHARSET utf8 COLLATE utf8_general_ci", conf.DbName)
		dropdbsql = fmt.Sprintf("DROP DATABASE IF EXISTS `%S`", conf.DbName)
		if force {
			fmt.Println(dropdbsql)
		}
		fmt.Println(createdbsql)
		break
	default:
		seelog.Critical("db dirve not support:", conf.DbType)
		return
	}
	db, err := sql.Open(conf.DbType, conf.DbName)
	if err != nil {
		panic(err.Error())
	}
	if force {
		_, err = db.Exec(dropdbsql)
	}
	_, err1 := db.Exec(createdbsql)
	if err != nil || err1 != nil {
		seelog.Error("db exec error:", err, err1)
		panic(err.Error())
	} else {
		seelog.Trace("database ", conf.DbName, " created")
	}
	defer db.Close()
	fmt.Println("create database end")
}

func TrunCateRbacTable(table []string) {
	seelog.Trace("delete table start")
	var dns, sqlstring string
	switch conf.DbType {
	case "sqlite3":
		dns = conf.DbName
	default:
		seelog.Trace("db drive not support:", conf.DbType)
	}
	db, err := sql.Open(conf.DbType, dns)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	for _, i := range table {
		seelog.Trace("table deleting:" + i)
		sqlstring = fmt.Sprintf("TRUNCATE TABLE　`%s`", i)
		_, err = db.Exec(sqlstring)
		if err != nil {
			seelog.Trace("table delete error:" + err.Error())
			panic(err.Error())
		} else {
			seelog.Trace("table delete success:" + i)
		}
	}
	seelog.Trace("delete table end")
}

func Connect() {
	
	RegisterDBModel()
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	err := orm.RegisterDataBase("default",_SQLITE3_DRIVER, conf.DbName, 10)
	if err != nil {
		seelog.Error("register data:" + err.Error())
		panic(err.Error())
	}
	if conf.DbLog == "open" {
		seelog.Trace("develop mpde, debug database: db.log")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			seelog.Error(e.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}
    
}
