package request

type CreateFavoriteRequest struct {
	UserID   string `json:"id" binding:"required"`
	TicketID string `json:"ticketId" binding:"required"`
}

type ListFavoriteRequest struct {
	PageNo   int    `json:"pageNo" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Order    string `json:"order"`
	OrderBy  string `json:"orderBy"`
	UserID   string `json:"userId"`
}

type DeleteFavoriteRequest struct {
	ID string `json:"id"`
}
