package userp

import "github.com/provider-go/userp/defaultuserp"

type UserProvider interface {
	Create(userId, belongId, username, password, nickname string, sex int, avatar, phone, email string) error
	Delete(userId string) error
	List(belongId, pageSize, pageNum string) (string, int64, error)
	View(userId string) (string, error)
	Login(belongId, username, password string) error
	ResetPassword(userId, password string) error
}

func GetUserProvider(name string) UserProvider {
	if name == "default" {
		return defaultuserp.NewDefaultUserProvider()
	}

	return nil
}
