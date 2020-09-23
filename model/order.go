package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Order struct {
	ID         string `gorm:"id"`
	TicketID   string `gorm:"ticket_id"`
	TicketName string
	UserID     string `gorm:"user_id"`
	UserName   string
	Status     string
	Price      float32
	Count      int
	CreateTime time.Time
	UpdateTime time.Time
}

func GetOrder(ID string) (Order, error) {
	var order Order

	err := config.DB.Table(enum.TABLENAME_ORDER).Where(Order{ID: ID}).Find(&order).Error
	return order, err
}

func CreateOrder(newOrder Order) (Order, error) {
	if err := config.DB.Table(enum.TABLENAME_ORDER).Create(newOrder).Error; err != nil {
		return Order{}, err
	}

	return newOrder, nil
}

func ListOrders(filter Order, p Pagination, orderf OrderFilter) ([]Order, int, error) {
	var orders []Order
	var totalCount int

	database := config.DB.Table(enum.TABLENAME_ORDER).Where(&filter)
	database.Count(&totalCount)

	database = database.Order(orderf.Field + " " + orderf.Direction)
	database = database.Offset(p.Offset).Limit(p.Size)
	database.Find(&orders)

	if database.Error != nil {
		return nil, 0, database.Error
	}
	return orders, totalCount, nil
}

func UpdateOrder(filter, newOrder Order) error {
	return config.DB.Table(enum.TABLENAME_ORDER).Where(&filter).Update(newOrder).Error
}

func DeleteOrder(ID string) error {
	return config.DB.Table(enum.TABLENAME_ORDER).Delete(&Order{ID: ID}).Error
}
