package models

import (
	"FurballCommunity_backend/config/database"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	UserID      uint    `gorm:"primary_key" json:"user_id"`
	Account     string  `json:"account"  gorm:"not null"`
	Password    string  `json:"password"  gorm:"not null"`
	Pets        []Pet   `gorm:"foreign_key:UserID"`
	Phone       string  `json:"phone"`
	Username    string  `json:"username"  default:"请输入用户名"`
	Authority   uint    `json:"authority"`
	Gender      uint    `json:"gender"`
	Address     string  `json:"address"`
	Score       float64 `json:"score"`
	Intro       string  `json:"intro"`
	IDNumber    string  `json:"id_number"`
	Avatar      string  `json:"avatar"`
	PetExp      string  `json:"pet_exp"`
	WorkTime    uint    `json:"work_time"`
	PetNum      int     `json:"pet_num"`
	IDCardPhoto string  `json:"id_card_photo"`
	LikedBlog   string  `json:"liked_blog" gorm:"-"`
}

// HasMany 在User模型中定义HasMany方法，表示一个User拥有多个Pet
func (user *User) HasMany() interface{} {
	return &[]Pet{}
}

// CreateUser 创建用户
func CreateUser(user *User) (err error) {
	err = database.DB.Create(&user).Error
	return
}

type UserBase struct {
	UserID   uint   `json:"user_id"`
	Account  string `json:"account"`
	Username string `json:"username"`
}

// 获取用户列表
func GetUserList() (userList []*User, err error) {
	if err = database.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id获取单个用户
func GetUserById(id uint) (user *User, err error) {
	user = new(User)
	if err = database.DB.Preload("Pets").Where("user_id = ?", id).First(user).Error; err != nil {
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

// 根据phone获取用户
func GetUserByPhone(phone string) (user *User, err error) {
	user = new(User)
	if err = database.DB.Where("phone = ?", phone).First(user).Error; err != nil {
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

func UpdateUserInfo(user *User) (err error) {
	err = database.DB.Model(&user).Updates(map[string]interface{}{
		"phone":         user.Phone,
		"username":      user.Username,
		"authority":     user.Authority,
		"gender":        user.Gender,
		"address":       user.Address,
		"score":         user.Score,
		"intro":         user.Intro,
		"id_number":     user.IDNumber,
		"avatar":        user.Avatar,
		"pet_exp":       user.PetExp,
		"work_time":     user.WorkTime,
		"pet_num":       user.PetNum,
		"ID_card_photo": user.IDCardPhoto,
	}).Error
	return
}

func convertStringSliceToUintSlice(strSlice []string) ([]uint, error) {
	uintSlice := make([]uint, 0, len(strSlice))
	for _, str := range strSlice {
		uintVal, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		uintSlice = append(uintSlice, uint(uintVal))
	}
	return uintSlice, nil
}

func GetUserLikedBlog(userID uint) (likedBlog []uint, err error) {
	var likedBlogStr string
	if err = database.DB.Where("user_id = ?", userID).First(likedBlogStr).Error; err != nil {
		return nil, err
	}
	likedBlogChar := strings.Split(likedBlogStr, ",")
	likedBlog, err = convertStringSliceToUintSlice(likedBlogChar)
	if err != nil {
		fmt.Println("转换失败：", err)
		return
	}
	return
}
