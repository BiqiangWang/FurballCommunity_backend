package models

import (
	"FurballCommunity_backend/config/database"
	"log"
	"time"
)

type Order struct {
	OrderID      uint       `gorm:"primary_key" json:"order_id"`
	PetID        uint       `json:"pet_id"`
	AnnouncerID  uint       `json:"announcer_id"`
	ReceiverID   uint       `json:"receiver_id"`
	Pet          Pet        `gorm:"foreign_key:PetID"`
	OrderCmts    []OrderCmt `gorm:"foreign_key:OrderID"`
	AnnounceTime time.Time  `json:"announce_time" gorm:"default:CURRENT_TIMESTAMP"`
	StartTime    time.Time  `json:"start_time" gorm:"default:CURRENT_TIMESTAMP"`
	EndTime      time.Time  `json:"end_time" gorm:"default:CURRENT_TIMESTAMP"`
	Place        string     `json:"place"`
	PetHealth    string     `json:"pet_health"`
	Status       int        `json:"status"`
	Remark       string     `json:"remark"`
	Price        int        `json:"price"`
	Evaluation   string     `json:"evaluation"`
	Score        float64    `json:"score"`
}

// BelongsTo 在Pet模型中定义BelongsTo方法，表示Order属于一个pet
func (order *Order) BelongsTo() interface{} {
	return &Pet{}
}

// HasMany 在Order模型中定义HasMany方法，表示一个Order拥有多个OrderCmt
func (order *Order) HasMany() interface{} {
	return &[]OrderCmt{}
}

func CreateOrder(order *Order) (err error) {
	err = database.DB.Create(&order).Error
	return
}

func GetOrderList(userID uint) (orderList []*Order, err error) {
	if err := database.DB.Where("announcer_id = ?", userID).Find(&orderList).Error; err != nil {
		return nil, err
	}
	return orderList, nil
}

func GetOrderOfPet(petID uint) (order []*Order, err error) {
	log.Printf("GetOrderOfpet: petID=%d\n", petID)
	if err := database.DB.Where("pet_id = ?", petID).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderInfoByID(orderID uint) (order *Order, err error) {
	if err := database.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func UpdateOrderInfo(order *Order) (err error) {
	err = database.DB.Model(&order).Updates(map[string]interface{}{
		"receiver_id": order.ReceiverID,
		"start_time":  order.StartTime,
		"end_time":    order.EndTime,
		"place":       order.Place,
		"pet_health":  order.PetHealth,
		"status":      order.Status,
		"remark":      order.Remark,
		"price":       order.Price,
		"evaluation":  order.Evaluation,
		"score":       order.Score,
	}).Error
	return
}

//func DeleteOrderOfPet(orders []*Order, petID uint) (err error) {
//	// 开始数据库事务
//	tx := database.DB.Begin()
//
//	// 删除所有订单
//	if err := tx.Where("pet_id = ?", petID).Delete(&orders).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//
//	return
//}

func DeleteOrder(orderID uint) (err error) {
	err = database.DB.Delete(&Order{}, orderID).Error
	return
}
