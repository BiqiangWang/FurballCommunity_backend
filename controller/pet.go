package controller

import (
	"FurballCommunity_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddPet
// @Summary 添加宠物
// @Description 添加一个新的宠物 eg：{"pet_name":"xiaohuang","user_id":2 }
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param   pet    body    string   true      "petname + userid"
// @Success 200 {string} string	"ok"
// @Router /v1/pet/add [post]
func AddPet(c *gin.Context) {
	var pet models.Pet
	err := c.BindJSON(&pet)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}

	if err := models.AddPet(&pet); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "添加成功", "pet_id": pet.PetID})
	}
}

// GetPetInfoByID
// @Summary 通过宠物id查询宠物信息
// @Description 通过宠物id查询宠物信息
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "pet_id"
// @Success 200 {string} string	"ok"
// @Router /v1/pet/getPetInfoByID/{id} [get]
func GetPetInfoByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id！"})
		return
	}
	petId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	pet, e := models.GetPetInfoByID(uint(petId))
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "查询成功", "pet_info": pet})
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
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	petID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}

	if err := models.DeleteOrderOfPet(uint(petID)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	}

	if err := models.DeletePet(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "删除成功",
		})
	}
}

// GetPetList
// @Summary 获取宠物列表
// @Description 根据用户id获取宠物列表
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Router /v1/pet/getPetList/{id} [get]
func GetPetList(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	petList, err := models.GetPetList(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功获取宠物列表", "pet_list": petList})

}

// UpdatePetInfo
// @Summary 更改宠物信息
// @Description 通过id，更新宠物信息，包括宠物名称、年龄、重量、绝育信息、品种和健康情况等 eg：{"pet_name":"wangwang", "gender":1, "age":2, "weight":33, "sterilization":1, "breed":"taidi", "health":"yes" }
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_pet_info"
// @Success 200 {string} string	"ok"
// @Router /v1/pet/updatePetInfo/{id} [put]
func UpdatePetInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	petId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}

	pet, err := models.GetPetInfoByID(uint(petId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}

	e := c.BindJSON(&pet)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}

	if err := models.UpdatePetInfo(pet); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"meg":  "成功修改宠物信息！",
			"info": pet,
		})
	}
}
