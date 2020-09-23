package request

type GetOrderRequest struct {
	ID string `json:"id"`
}

type CreateOrderRequest struct {
	TicketID string  `json:"ticketId" binding:"required"`
	UserID   string  `json:"userId" binding:"required"`
	Price    float32 `json:"price"`
	Count int `json:"count" binding:"required"`
}

type ListOrderRequest struct {
	PageNo   int    `json:"pageNo" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	Status   string `json:"status"`
	UserID   string `json:"userId"`
}

type UpdateOrderRequest struct {
	ID     string `json:"id" binding:"required"`
	Status string `json:"status"`
}

type DeleteOrderRequest struct {
	ID string `json:"id" binding:"required"`
}

type PayForOrderRequest struct {
	ZjpayID string `json:"zjpayId" binding:"required"`
	OrderID string `json:"orderId" binding:"required"`
}

type ListFinishedTicketRequest struct {
	PageNo    int    `json:"pageNo" binding:"required"`
	PageSize  int    `json:"pageSize" binding:"required"`
	OrderBy   string `json:"orderBy"`
	Order     string `json:"order"`
	UserID    string `json:"userId"`
	OutOfDate string `json:"outOfDate" binding:"required"`
}
