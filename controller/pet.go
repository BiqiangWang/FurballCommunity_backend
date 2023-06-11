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
// @Tags Pet
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

// GetPetInfoByID
// @Summary 通过宠物id查询宠物信息
// @Description 通过宠物id查询宠物信息 eg：{ "pet_id":2 }
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Success 200 {string} string	"ok"
// @Router /v1/pet/getPetInfoByID [GET]
func GetPetInfoByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id！"})
		return
	}
	pet, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&pet)
	c.JSON(http.StatusOK, gin.H{"pet_info": pet})
}

// DeletePet
// @Summary 删除宠物
// @Description 通过id，删除宠物 eg：{ "id":"5"}
// @Tags Pet
// @Accept  json
// @Param   id    path    uint     true      "id"
// @Router /v1/pet/deletePet/{id} [delete]
func DeletePet(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeletePet(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}
