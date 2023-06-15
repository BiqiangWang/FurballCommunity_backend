package models

import (
	"FurballCommunity_backend/config/database"
	"gorm.io/gorm"
	"log"
)

type Pet struct {
	PetID         uint    `gorm:"primary_key" json:"pet_id"`
	UserID        uint    `json:"user_id" gorm:"not null"`
	PetName       string  `json:"pet_name"`
	Orders        []Order `gorm:"foreign_key:PetID"`
	Gender        int     `json:"gender"`
	Age           int     `json:"age"`
	Weight        int     `json:"weight"`
	Sterilization int     `json:"sterilization"`
	Breed         string  `json:"breed"`
	Health        string  `json:"health"`
	// photo entry have not been added in this table
}

// AddPet means add a pet to pet table
func AddPet(pet *Pet) (err error) {
	err = database.DB.Create(&pet).Error
	return
}

// GetPetList
// 基于预加载方式的联表查询
// 根据用户id获取该用户的宠物列表
func GetPetList(userID uint) (petList []*Pet, err error) {
	log.Printf("GetPetList: userID=%d\n", userID)
	if err := database.DB.Where("user_id = ?", userID).Find(&petList).Error; err != nil {
		return nil, err
	}
	return petList, nil
}

// UpdatePetInfo
// 更新宠物信息，包括宠物名称、年龄、重量、绝育信息、品种和健康情况等
func UpdatePetInfo(pet *Pet) (err error) {
	err = database.DB.Model(&pet).Updates(map[string]interface{}{
		"pet_name":      pet.PetName,
		"gender":        pet.Gender,
		"age":           pet.Age,
		"weight":        pet.Weight,
		"sterilization": pet.Sterilization,
		"breed":         pet.Breed,
		"health":        pet.Health,
	}).Error
	return
}

// GetPetInfoByID
// 通过宠物id获取宠物信息
func GetPetInfoByID(petID uint) (pet *Pet, err error) {
	pet = new(Pet)
	if err = database.DB.Where("pet_id = ?", petID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id", "account", "username")
	}).First(pet).Error; err != nil {
		return nil, err
	}
	return
}

// DeletePet  删除宠物
func DeletePet(petID string) (err error) {
	err = database.DB.Delete(&Pet{}, petID).Error
	return
}

func DeleteOrderOfPet(petID uint) (err error) {
	// 开始数据库事务
	tx := database.DB.Begin()

	// 查询是否有对应的订单
	var orderCount int64
	if err := tx.Model(&Order{}).Where("pet_id = ?", petID).Count(&orderCount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if orderCount > 0 {
		// 如果有对应的订单，则先删除订单
		if err := tx.Where("pet_id = ?", petID).Delete(&Order{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	tx.Commit()

	return nil
}

// GetPetInfoByName
// 通过宠物名称获取宠物信息
//func GetPetInfoByName(petName string) (pet *Pet, err error) {
//	pet = new(Pet)
//	if err = database.DB.Where("pet_name = ?", petName).Preload("User", func(db *gorm.DB) *gorm.DB {
//		return db.Select("user_id", "username", "account")
//	}).First(pet).Error; err != nil {
//		return nil, err
//	}
//	return
//}
