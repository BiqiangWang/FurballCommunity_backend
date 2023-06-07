package controller

import (
	"FurballCommunity_backend/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// AddPet
// @Summary 添加宠物
// @Description 添加一个新的宠物 eg：{ "pet_name":"xiaohuang", "user_id":2 }
// @Accept  json
// @Produce  json
// @Param   pet    body    string   true      "petname + userid"
// @Success 200 {string} string	"ok"
// @Router /v1/pet/add [post]
func AddPet(c *gin.Context) {
	var pet models.Pet
	c.BindJSON(&pet)

	_, err := models.GetPetInfoByName(pet.PetName)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 宠物名已存在
		c.JSON(http.StatusOK, gin.H{"status": reStatusError, "text": "该宠物已被添加"})
	} else {
		if err := models.AddPet(&pet); err != nil {
			c.JSON(http.StatusCreated, gin.H{"status": reStatusError, "text": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": reStatusSuccess, "text": "添加成功", "petid": pet.PetID})
		}
	}
}
