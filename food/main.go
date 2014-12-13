package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"

	//orm
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30

	orm.RegisterDataBase("default", "mysql", "test:test@/go?charset=utf8", maxIdle, maxConn)

	orm.DefaultTimeLoc = time.UTC

}

func main() {

	//orm_handler.Using("default") // 默认使用 default，你可以指定为其他数据库

	beego.Run()
}
