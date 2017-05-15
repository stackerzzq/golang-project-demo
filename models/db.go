package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.Debug, _ = beego.AppConfig.Bool("DB::debug")

	switch beego.AppConfig.String("DB::driver") {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s",
			beego.AppConfig.String("DB::user"),
			beego.AppConfig.String("DB::pass"),
			beego.AppConfig.String("DB::host"),
			beego.AppConfig.String("DB::port"),
			beego.AppConfig.String("DB::name"),
			"Asia%2FShanghai"), 10)
	case "sqlite":
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file"), 10)
	default:
		fmt.Println("数据库配置错误")
	}

	orm.RunSyncdb("default", false, true)
}
