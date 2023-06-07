package models

import "FurballCommunity_backend/config/database"

type Pet struct {
	PetID         uint    `gorm:"primary_key" json:"pet_id"`
	UserID        uint    `json:"user_id" binding:"required"`
	PetName       string  `json:"pet_name" binding:"required"`
	Gender        int     `json:"gender"`
	Age           int     `json:"age"`
	Weight        float64 `json:"weight"`
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
	if err := database.DB.Preload("Pet").Where("user_id = ?", userID).Find(&petList).Error; err != nil {
		return nil, err
	}
	return petList, nil
}

// UpdatePetInfo
// 更新宠物信息，包括宠物名称、年龄、重量、绝育信息、品种和健康情况等
func UpdatePetInfo(pet *Pet) (err error) {
	err = database.DB.Model(&pet).Updates(map[string]interface{}{
		"gender":        pet.Gender,
		"pet_name":      pet.PetName,
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
	if err = database.DB.Where("pet_id = ?", petID).First(pet).Error; err != nil {
		return nil, err
	}
	return
}

// GetPetInfoByName
// 通过宠物名称获取宠物信息
func GetPetInfoByName(petName string) (pet *Pet, err error) {
	pet = new(Pet)
	if err = database.DB.Where("pet_name = ?", petName).First(pet).Error; err != nil {
		return nil, err
	}
	return
}

// DeletePet  删除宠物
func DeletePet(petID uint) (err error) {
	err = database.DB.Delete(&Pet{}, petID).Error
	return
}
