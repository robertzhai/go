package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Food struct {
	Id   int
	Dish string `orm:"size(100)"`
	Pic  string `orm:"size(300)"`
}

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
	orm.RegisterModel(new(Food))
}

func (this *Food) Query() ([]Food, error) {
	o := orm.NewOrm()
	var data []Food
	_, err := o.QueryTable("food").Limit(12, 0).All(&data)
	return data, err
}
