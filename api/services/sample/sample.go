package sample

import (
	"github.com/abalol/beego-boilerplate/api/models"
	"github.com/astaxie/beego/orm"
)

// GetFirstSample 最新Ping取得
// description Pingの最初の1件を返す
func GetFirstSample() []models.Ping {
	o := orm.NewOrm()
	ping := []models.Ping{}
	o.QueryTable("ping").Filter("id", 1).All(&ping)
	return ping
}
