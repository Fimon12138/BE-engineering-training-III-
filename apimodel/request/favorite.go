package request

type CreateFavoriteRequest struct {
	UserID   string `json:"id"`
	TicketID string `json:"ticketId"`
}

type ListFavoriteRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Order    string `json:"order"`
	OrderBy  string `json:"orderBy"`
	UserID   string `json:"userId"`
}

type DeleteFavoriteRequest struct {
	ID string `json:"id"`
}
