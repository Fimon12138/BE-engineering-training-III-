package request

type GetTicketRequest struct {
	ID string `json:"id"`
}

type CreateTicketRequest struct {
	Name        string  `json:"name"`
	MerchantID  string  `json:"merchantId"`
	ImageRow    string  `json:"imageRow"`
	ImageColumn string  `json:"imageColumn"`
	ImageDetail string  `json:"imageDetail"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	StartTime   string  `json:"startTime"`
	Location    string  `json:"location"`
	Count       int     `json:"count"`
	Description string  `json:"Description"`
}

type UpdateTicketRequest struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	Price             float32 `json:"price"`
	Count             int     `json:"count"`
	SubscribeIncrease bool    `json:"subscribeCount"`
}

type ListTicketRequest struct {
	PageNo     int    `json:"pageNo" binding:"required"`
	PageSize   int    `json:"pageSize" binding:"required"`
	OrderBy    string `json:"orderBy"`
	Order      string `json:"order"`
	NameFilter string `json:"nameFilter"`
	Type       string `json:"type"`
}

type DeleteTicketRequest struct {
	ID string `json:"id"`
}
