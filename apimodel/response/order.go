package response

import (
	"tickethub_service/model"
	"tickethub_service/util"
)

type Order struct {
	ID         string  `json:"id"`
	TicketID   string  `json:"ticketId"`
	TicketName string  `json:"ticketName"`
	UserID     string  `json:"userId"`
	UserName   string  `json:"userName"`
	Status     string  `json:"status"`
	Price      float32 `json:"price"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
}

type OrderWithTicket struct {
	Order
	Ticket
}

type ListOrderWithTicket struct {
	PageNo     int     `json:"pageNo"`
	PageSize   int     `json:"pageSize"`
	Result     []OrderWithTicket `json:"result"`
	TotalCount int     `json:"totalCount"`
}

type GetOrderResponse struct {
	Order
}

type CreateOrderResponse struct {
	Order
}

type ListOrderResponse struct {
	PageNo     int     `json:"pageNo"`
	PageSize   int     `json:"pageSize"`
	Result     []Order `json:"result"`
	TotalCount int     `json:"totalCount"`
}

func (o *Order) Load(order model.Order) {
	o.ID = order.ID
	o.TicketID = order.TicketID
	o.TicketName = order.TicketName
	o.UserID = order.UserID
	o.UserName = order.TicketName
	o.Status = order.Status
	o.Price = order.Price
	o.CreateTime = util.GetDisplayTime(order.CreateTime)
	o.UpdateTime = util.GetDisplayTime(order.UpdateTime)
}

type ListFinishedTicketResponse struct {
	PageNo     int      `json:"pageNo"`
	PageSize   int      `json:"pageSize"`
	Result     []Ticket `json:"result"`
	TotalCount int      `json:"totalCount"`
}
