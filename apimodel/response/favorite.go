package response

type CreateFavoriteResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	TicketID string `json:"ticketId"`
}


type ListFavoriteResponse struct {
	PageNo     int               `json:"pageNo"`
	PageSize   int               `json:"pageSize"`
	Result     []Ticket `json:"result"`
	TotalCount int               `json:"totalCount"`
}

