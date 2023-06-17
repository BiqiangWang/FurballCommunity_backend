package models

import (
	"FurballCommunity_backend/config/database"
	"log"
	"time"
)

type OrderCmt struct {
	OrderCmtID  uint      `gorm:"primary_key" json:"order_cmt_id"`
	OrderID     uint      `json:"order_id"`
	UserID      uint      `json:"user_id"`
	CommentTime time.Time `json:"comment_time" gorm:"default:CURRENT_TIMESTAMP"`
	Content     string    `json:"content"`
	Order       Order     `foreign_key:"OrderID"`
	IsRoot      bool      `json:"is_root"`
	ParentID    uint      `json:"parent_id"'`
	//ChildCmts   []OrderCmt `gorm:"foreignKey:OrderCmtID;constraint:OnDelete:CASCADE"`
}

// BelongsTo 在Pet模型中定义BelongsTo方法，表示OrderCmt属于一个order
func (orderCmt *OrderCmt) BelongsTo() interface{} {
	return &Order{}
}

// HasMany 在orderCmt模型中定义HasMany方法，表示一个cmt可以拥有多个child_cmt
func (orderCmt *OrderCmt) HasMany() interface{} {
	return &[]OrderCmt{}
}

func CreateOrderCmt(orderCmt *OrderCmt) (err error) {
	err = database.DB.Create(&orderCmt).Error
	return
}

func GetCmtListOfOrder(orderID uint) (orderCmt []*OrderCmt, err error) {
	log.Printf("GetCmtListOfOrder: orderID=%d\n", orderID)
	if err := database.DB.Where("order_id = ?", orderID).Find(&orderCmt).Error; err != nil {
		return nil, err
	}
	return orderCmt, nil
}

func DeleteOrderCmt(orderCmtID uint) (err error) {
	err = database.DB.Delete(&OrderCmt{}, orderCmtID).Error
	return
}
