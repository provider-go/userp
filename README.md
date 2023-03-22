# userp

## UserProvider 模块接口
```
type UserProvider interface {
	// SetDB 设置数据库连接
	SetDB(db *gorm.DB)
	// Create 创建用户
	Create(userId, belongId, username, password, nickname string, sex int, avatar, phone, email string) error
	// Delete 删除用户
	Delete(userId string) error
	// List 获取用户列表
	List(belongId, pageSize, pageNum string) (string, int64, error)
	// View 查看用户详情
	View(userId string) (string, error)
	// Login 用户登录
	Login(belongId, username, password string) error
	// ResetPassword 重置密码
	ResetPassword(userId, password string) error
}
```

## UserProvider 使用示例
```
package main

import (
	"fmt"
	"github.com/provider-go/userp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	DBConn, err := gorm.Open(mysql.Open("root:pawword@tcp(ip:3306)/db?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	provider := userp.GetUserProvider("default")
	provider.SetDB(DBConn)
	res, err := provider.View("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```