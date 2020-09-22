package response

type ListRecommendationResponse struct {
	Tickets []Ticket `json:"ticket_ids"`
}
