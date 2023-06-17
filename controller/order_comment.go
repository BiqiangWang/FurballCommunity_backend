package controller

import (
	"FurballCommunity_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrderComment
// @Summary 创建订单评论
// @Description 根据用户id、订单id，添加一个新的订单评论 eg：{ "user_id":1, "order_id":2, "content":"宠物照顾的针布戳呢" }
// @Tags OrderCmt
// @Accept  json
// @Produce  json
// @Param   orderCmt    body    string   true      "userid + orderId + content"
// @Success 200 {string} string	"ok"
// @Router /v1/orderCmt/create [post]
func CreateOrderComment(c *gin.Context) {
	var orderCmt models.OrderCmt
	c.BindJSON(&orderCmt)

	if err := models.CreateOrderCmt(&orderCmt); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功添加订单评论", "orderCmt_info": orderCmt})
	}
}

// GetCommentListOfOrder
// @Summary 获取订单的评论列表
// @Description 根据订单id获取评论列表
// @Tags OrderCmt
// @Accept  json
// @Produce  json
// @Param   order_id    path    uint     true      "order_id"
// @Router /v1/orderCmt/getOrderCmtList/{order_id} [get]
func GetCommentListOfOrder(c *gin.Context) {
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
	cmtList, err := models.GetCmtListOfOrder(uint(orderId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功获取订单评论列表", "orderCmt_list": cmtList})
}

// DeleteOrderCmt
// @Summary 删除订单评论
// @Description 通过orderCmtID，删除宠物 eg：{ "orderCmtID":"5"}
// @Tags OrderCmt
// @Accept  json
// @Param   order_cmt_id    path    uint     true      "order_cmt_id"
// @Router /v1/orderCmt/deleteOrderCmt/{order_cmt_id} [delete]
func DeleteOrderCmt(c *gin.Context) {
	id, ok := c.Params.Get("order_cmt_id}")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	orderCmtID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	if err := models.DeleteOrderCmt(uint(orderCmtID)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "订单评论删除成功",
		})
	}
}
