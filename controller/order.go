package controller

import (
	"FurballCommunity_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrder
// @Summary 创建订单
// @Description 根据用户id，创建一个新的订单 eg：{ "pet_id":3, "announcer_id":2 }
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order    body    string   true      "petname + userid"
// @Success 200 {string} string	"ok"
// @Router /v1/order/create [post]
func CreateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)

	if err := models.CreateOrder(&order); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {

		c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功创建订单", "order_info": order})
	}
}

// GetOrderList
// @Summary 获取用户的订单列表
// @Description 根据用户id获取订单列表
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   user_id    path    uint     true      "user_id"
// @Router /v1/order/getOrderList/{user_id} [get]
func GetOrderList(c *gin.Context) {
	id, ok := c.Params.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	orderList, err := models.GetOrderList(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功获取订单列表", "order_list": orderList})
}

// GetOrderOfPet
// @Summary 获取宠物的订单
// @Description 根据宠物id获取订单
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   pet_id    path    uint     true      "pet_id"
// @Router /v1/order/getOrderOfPet/{pet_id} [get]
func GetOrderOfPet(c *gin.Context) {
	id, ok := c.Params.Get("pet_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	petId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	order, err := models.GetOrderOfPet(uint(petId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功宠物订单", "order": order})
}

// UpdateOrderInfo
// @Summary 更改订单信息
// @Description 通过订单id，更新接收者、开始结束时间、地点、健康、订单状态、备注、价格、评价、评分等 eg：{"receiver_id":1}
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order_id    path    uint     true      "order_id"
// @Param   user    body    string     true      "new_order_info"
// @Success 200 {string} string	"ok"
// @Router /v1/order/updateOrderInfo/{order_id} [put]
func UpdateOrderInfo(c *gin.Context) {
	id, ok := c.Params.Get("order_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	order, err := models.GetOrderInfoByID(uint(orderId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	e := c.BindJSON(&order)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}

	if err := models.UpdateOrderInfo(order); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"meg":  "成功修改订单信息！",
			"info": order,
		})
	}
}

// GetOrderInfoById
// @Summary 获取订单信息
// @Description 根据订单id获取详情
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order_id    path    uint     true      "order_id"
// @Router /v1/order/getOrderInfoById/{order_id} [get]
func GetOrderInfoById(c *gin.Context) {
	id, ok := c.Params.Get("order_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	order, err := models.GetOrderInfoByID(uint(orderId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功获取订单信息", "order": order})
}

// DeleteOrder
// @Summary 删除订单
// @Description 根据订单id删除订单
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order_id    path    uint     true      "order_id"
// @Router /v1/order/delete/{order_id} [delete]
func DeleteOrder(c *gin.Context) {
	id, ok := c.Params.Get("order_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	if err := models.DeleteOrder(uint(orderId)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "删除成功",
		})
	}
}
