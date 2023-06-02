package models

import (
	"FurballCommunity_backend/config/database"
)

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Account   string `json:"account"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	Username  string `json:"username"  default:"请输入用户名"`
	Authority uint   `json:"authority"`
}

// 创建用户
func CreateUser(user *User) (err error) {
	err = database.DB.Create(&user).Error
	return
}

// 获取用户列表
func GetUserList() (userList []*User, err error) {
	if err = database.DB.Select("id", "account", "username").Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id获取单个用户
func GetUserById(id string) (user *User, err error) {
	user = new(User)
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// 根据username获取用户
func GetUserByAccount(account string) (user *User, err error) {
	user = new(User)
	if err = database.DB.Where("account = ?", account).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// 更新用户名
func UpdateUserName(user *User) (err error) {
	err = database.DB.Select("username").Updates(user).Error
	return
}

// 更新用户密码
func UpdatePassword(user *User) (err error) {
	err = database.DB.Select("password").Updates(user).Error
	return
}

// 删除用户
func DeleteUser(id string) (err error) {
	err = database.DB.Delete(&User{}, id).Error
	return
}
