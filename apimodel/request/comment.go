package request

type CreateCommentRequest struct {
	UserID   string `json:"userId" binding:"required"`
	TicketID string `json:"ticketId" binding:"required"`
	Content  string `json:"content"`
}

type ListCommentRequest struct {
	PageNo   int    `json:"pageNo" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	UserID   string `json:"userId"`
	TicketID string `json:"ticketId"`
}

type UpdateCommentRequest struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type DeleteCommentRequest struct {
	ID string `json:"id"`
}
