package ping

import (
	"app/api/models"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// UpdateOnePings PingUpdate
// description Pingの最初の1件を現在日時で更新する
func UpdateOnePings() models.Ping {
	o := orm.NewOrm()

	// 現在時刻に更新してみる
	n := time.Now()

	ping := models.Ping{Id: 1}
	if o.Read(&ping) == nil {
		ping.Content = n.Format("2006-01-02")
		if num, err := o.Update(&ping); err == nil {
			fmt.Println(num)
		}
	}
	return ping
}
