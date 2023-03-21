package userp

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func connMysql() UserProvider {
	// 连接mysql
	DBConn, err := gorm.Open(mysql.Open("root:******@tcp(ip:13306)/services?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	provider := GetUserProvider("default")
	provider.SetDB(DBConn)
	return provider
}

func TestCreate(t *testing.T) {
	provider := connMysql()
	err := provider.Create("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "bbbbbbb", "yuyu", "12346", "yuyu", 1, "HHHHH", "15101131912", "69401295@qq.com")
	if err != nil {
		t.Log(err)
	}
}

func TestView(t *testing.T) {
	provider := connMysql()
	res, err := provider.View("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestList(t *testing.T) {
	provider := connMysql()
	res, count, err := provider.List("bbbbbbb", "20", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(count)
	t.Log(res)
}

func TestLogin(t *testing.T) {
	provider := connMysql()
	err := provider.Login("bbbbbbb", "qiqi", "12346")
	if err != nil {
		t.Log(err)
	}
}

func TestResetPassword(t *testing.T) {
	provider := connMysql()
	err := provider.ResetPassword("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1234678")
	if err != nil {
		t.Log(err)
	}
}

func TestDelete(t *testing.T) {
	provider := connMysql()
	err := provider.Delete("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	if err != nil {
		t.Log(err)
	}
}
