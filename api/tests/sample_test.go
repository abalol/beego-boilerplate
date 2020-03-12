package test

import (
	_ "app/api/controllers"
	"app/api/models"
	_ "app/api/routers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
	orm.RegisterDataBase("default", "mysql", datasource, 30)
	orm.RunSyncdb("default", false, true)
	beego.TestBeegoInit("/go/src/app/api/tests")
}

// TestGet コントローラのテスト
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/ping/sample", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	// Assertions
	var ping models.Ping
	json.Unmarshal([]byte(w.Body.String()), &ping)

	// Equal 項目単位でみたい時
	assert.Equal(t, w.Code, 200, "コードが正常（２００）であること")
	assert.Equal(t, ping.Id, int64(1), "idが1であること")
	assert.Equal(t, ping.Content, "2020-03-12", "Contentの値が正常であること")

	// Objectでの比較。DBの値をそのまま吐くようなエンドポイントのテストならこっちが楽。
	var ex models.Ping
	ex.Id = 1
	ex.Content = "2020-03-12"
	assert.Exactly(t, ping, ex, "Ping Objectでの比較")
}
