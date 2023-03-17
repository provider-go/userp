package defaultuserp

import (
	"github.com/provider-go/userp/defaultuserp/models/mysql"
	"testing"
)

func connMysql() {
	// 连接mysql
	mysql.NewMysql("root:ZTQ4ZTBjMTViNGMzODgzODUz@tcp(120.53.243.73:13306)/services?charset=utf8")
}

func TestCreate(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	err := provider.Create("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "bbbbbbb", "yuyu", "12346", "yuyu", 1, "HHHHH", "15101131912", "69401295@qq.com")
	if err != nil {
		t.Log(err)
	}
}

func TestView(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	res, err := provider.View("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestList(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	res, count, err := provider.List("bbbbbbb", "20", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(count)
	t.Log(res)
}

func TestLogin(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	err := provider.Login("bbbbbbb", "qiqi", "12346")
	if err != nil {
		t.Log(err)
	}
}

func TestResetPassword(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	err := provider.ResetPassword("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "1234678")
	if err != nil {
		t.Log(err)
	}
}

func TestDelete(t *testing.T) {
	connMysql()
	provider := NewDefaultUserProvider()
	err := provider.Delete("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	if err != nil {
		t.Log(err)
	}
}
