package request

type GetOrderRequest struct {
	ID string `json:"id"`
}

type CreateOrderRequest struct {
	TicketID string  `json:"ticketId"`
	UserID   string  `json:"userId"`
	Price    float32 `json:"price"`
}

type ListOrderRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	Status   string `json:"status"`
	UserID   string `json:"userId"`
}

type UpdateOrderRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type PayForOrderRequest struct {
	ZjpayID string `json:"zjpayId"`
	OrderID string `json:"orderId"`
}

type ListFinishedTicketRequest struct {
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	OrderBy   string `json:"orderBy"`
	Order     string `json:"order"`
	UserID    string `json:"userId"`
	OutOfDate string `json:"outOfDate"`
}
