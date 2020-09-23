package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Ticket struct {
	ID             string `gorm:"id"`
	Name           string
	MerchantID     string `grom:"merchant_id"`
	ImageRow       string
	ImageColumn    string
	ImageDetail    string
	Type           string
	Price          float32
	StartTime      time.Time
	Location       string
	Count          int
	SubscribeCount int
	Description    string
	CreateTime     time.Time
	UpdateTime     time.Time
}

func GetTicket(ID string) (Ticket, error) {
	var ticket Ticket

	err := config.DB.Table(enum.TABLENAME_TICKET).Where(&Ticket{ID: ID}).Find(&ticket).Error

	return ticket, err
}

func CreateTicket(newTicket Ticket) (Ticket, error) {
	if err := config.DB.Table(enum.TABLENAME_TICKET).Create(newTicket).Error; err != nil {
		return Ticket{}, err
	}
	return newTicket, nil
}

func ListTicket(filter Ticket, p Pagination, orderf OrderFilter, nameFilter string) ([]Ticket, int, error) {
	var tickets []Ticket
	var totalCount int

	database := config.DB.Table(enum.TABLENAME_TICKET).Where(&filter)
	database = database.Where(TICKET_FIELD_NAME+MYSQL_CONTANINS_STRING, "%"+nameFilter+"%")
	database.Count(&totalCount)

	database = database.Order(orderf.Field + " " + orderf.Direction)
	database = database.Offset(p.Offset).Limit(p.Size)
	database.Find(&tickets)

	if database.Error != nil {
		return nil, 0, database.Error
	}
	return tickets, totalCount, nil
}

func UpdateTicket(newTicket Ticket) error {
	return config.DB.Table(enum.TABLENAME_TICKET).Where(&Ticket{ID: newTicket.ID}).Update(newTicket).Error
}

func DeleteTicket(ID string) error {
	return config.DB.Table(enum.TABLENAME_TICKET).Delete(&Ticket{ID: ID}).Error
}
