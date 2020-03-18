package ping

import (
	"github.com/abalol/beego-boilerplate/models"
	"github.com/astaxie/beego/orm"
)

// GetFirstPing 最新Ping取得
// description Pingの最初の1件を返す
func GetFirstPing() []models.Ping {
	o := orm.NewOrm()
	ping := []models.Ping{}
	o.QueryTable("ping").All(&ping)
	return ping
}
