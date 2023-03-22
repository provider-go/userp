package userp

import (
	"github.com/provider-go/userp/defaultuserp"
	"gorm.io/gorm"
)

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

// GetUserProvider 获取用户模块供应商
func GetUserProvider(name string) UserProvider {
	if name == "default" {
		return defaultuserp.NewDefaultUserProvider()
	}

	return nil
}
