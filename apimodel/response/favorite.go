package response

import (
	"tickethub_service/model"
	"tickethub_service/util"
)

type CreateFavoriteResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	TicketID string `json:"ticketId"`
}

type DisplayFavorite struct {
	ID          string  `json:"id"`
	ImageColumn string  `json:"imageColumn"`
	ImageDetail string  `json:"imageDetail"`
	Price       float32 `json:"price"`
	StartTime   string  `json:"startTime"`
	Location    string  `json:"location"`
	Count       int     `json:"count"`
}

type ListFavoriteResponse struct {
	PageNo     int               `json:"pageNo"`
	PageSize   int               `json:"pageSize"`
	Result     []DisplayFavorite `json:"result"`
	TotalCount int               `json:"totalCount"`
}

func (d *DisplayFavorite) Load(ticket model.Ticket, favorite model.Favorite) {
	d.ID = favorite.ID
	d.ImageColumn = ticket.ImageColumn
	d.ImageDetail = ticket.ImageDetail
	d.Price = ticket.Price
	d.StartTime = util.GetDisplayTime(ticket.StartTime)
	d.Location = ticket.Location
	d.Count = ticket.Count
}
