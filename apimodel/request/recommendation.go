package request

type CreateRecommendationRequest struct {
	Location int    `json:"location"`
	TicketID string `json:"ticketId"`
}

type DeleteRecommendationRequest struct {
	Location int
}
