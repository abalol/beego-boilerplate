package main

import (
	_ "app/api/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
	orm.RegisterDataBase("default", "mysql", datasource, 30)
	orm.RunSyncdb("default", false, true)
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
