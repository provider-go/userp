package defaultuserp

import (
	"encoding/json"
	"errors"
	"github.com/provider-go/userp/defaultuserp/models/mysql"
	"strconv"
)

type DefaultUserProvider struct {
}

func NewDefaultUserProvider() *DefaultUserProvider {
	p := &DefaultUserProvider{}
	return p
}

func (d DefaultUserProvider) Create(userId, belongId, username, password, nickname string, sex int, avatar, phone, email string) error {
	return mysql.DBConn.Table("userp_default_user").Select("id", "belong_id", "username", "password", "nickname", "sex", "avatar", "phone", "email").
		Create(&mysql.UserPDefaultUser{Id: userId, BelongId: belongId, Username: username, Password: password, Nickname: nickname, Sex: sex, Avatar: avatar, Phone: phone, Email: email}).Error
}

func (d DefaultUserProvider) Delete(userId string) error {
	return mysql.DBConn.Debug().Table("userp_default_user").Where("id = ?", userId).Delete(&mysql.UserPDefaultUser{}).Error
}

func (d DefaultUserProvider) List(belongId, pageSize, pageNum string) (string, int64, error) {
	var rows []*mysql.UserPDefaultUser
	//计算列表数量
	var count int64
	mysql.DBConn.Table("userp_default_user").Where("belong_id = ?", belongId).Count(&count)
	// 获取分页数据
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageNumInt, _ := strconv.Atoi(pageNum)
	if err := mysql.DBConn.Table("userp_default_user").Where("belong_id = ?", belongId).Order("create_time desc").Limit(pageSizeInt).Offset((pageNumInt - 1) * pageSizeInt).Find(&rows).Error; err != nil {
		return "", 0, err
	}
	rowsByte, _ := json.Marshal(rows)
	return string(rowsByte), count, nil
}

func (d DefaultUserProvider) View(userId string) (string, error) {
	row := new(mysql.UserPDefaultUser)
	if err := mysql.DBConn.Table("userp_default_user").Where("id = ?", userId).First(&row).Error; err != nil {
		return "", err
	}
	rowByte, _ := json.Marshal(row)
	return string(rowByte), nil
}

func (d DefaultUserProvider) Login(belongId, username, password string) error {
	row := new(mysql.UserPDefaultUser)
	if err := mysql.DBConn.Table("userp_default_user").Where("belong_id = ? and username = ? and password = ?", belongId, username, password).First(&row).Error; err != nil {
		return err
	}
	if len(row.Id) != 32 {
		return errors.New("login fail")
	}
	return nil
}

func (d DefaultUserProvider) ResetPassword(userId, password string) error {
	return mysql.DBConn.Table("userp_default_user").Where("id = ?", userId).Updates(map[string]interface{}{
		"password": password,
	}).Error
}
