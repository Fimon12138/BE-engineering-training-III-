package request

type CreateCommentRequest struct {
	UserID   string `json:"userId"`
	TicketID string `json:"ticketId"`
	Content  string `json:"content"`
}

type ListCommentRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	UserID   string `json:"userId"`
	TicketID string `json:"ticketId"`
}

type UpdateCommentRequest struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	ID string `json:"id"`
}
