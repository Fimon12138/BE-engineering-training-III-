package response

import (
	model2 "tickethub_service/model"
	"tickethub_service/util"
)

type Ticket struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	MerchantID     string  `json:"merchantId"`
	ImageRow       string  `json:"imageRow"`
	ImageColumn    string  `json:"imageColumn"`
	ImageDetail    string  `json:"imageDetail"`
	Type           string  ` json:"type"`
	Price          float32 `json:"price"`
	StartTime      string  `json:"startTime"`
	Location       string  `json:"location"`
	Count          int     `json:"count"`
	SubscribeCount int     `json:"subscribeCount"`
	Description    string  `json:"description"`
	IsSubscribed string `json:"isSubscribed"`
	CreateTime     string  `json:"createTime"`
	UpdateTime     string  `json:"updateTime"`
}

type GetTicketResponse struct {
	Ticket
}

type CreateTicketResponse struct {
	Ticket
}

type ListTicketResponse struct {
	PageNo     int      `json:"pageNo"`
	PageSize   int      `json:"pageSize"`
	Result     []Ticket `json:"result"`
	TotalCount int      `json:"totalCount"`
}

func (t *Ticket) Load(ticket model2.Ticket) {
	t.ID = ticket.ID
	t.Name = ticket.Name
	t.MerchantID = ticket.MerchantID
	t.ImageRow = ticket.ImageRow
	t.ImageColumn = ticket.ImageColumn
	t.ImageDetail = ticket.ImageDetail
	t.Type = ticket.Type
	t.Price = ticket.Price
	t.StartTime = util.GetDisplayTime(ticket.StartTime)
	t.Location = ticket.Location
	t.Count = ticket.Count
	t.SubscribeCount = ticket.Count
	t.Description = ticket.Description
	t.CreateTime = util.GetDisplayTime(ticket.CreateTime)
	t.UpdateTime = util.GetDisplayTime(ticket.UpdateTime)
}
