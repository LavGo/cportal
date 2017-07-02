package main

import (
	_ "github.com/LavGo/cportal/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "cportal:cportal@/cportal?charset=utf8")
}
func main() {
	beego.Run()
}

