package service

import (
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/response"
	"tickethub_service/model"
	"tickethub_service/util/log"
	"time"
)

func CreateRecommendation(req request.CreateRecommendationRequest) error {
	recommendation := model.Recommendation{
		TicketID:   req.TicketID,
		Location:   req.Location,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := model.CreateRecommendation(recommendation); err != nil {
		log.Errorf("Failed to create recommendation[%v]: %v", recommendation, err)
		return err
	}

	return nil
}

func DeleteRecommendation(req request.DeleteRecommendationRequest) error {
	if err := model.DeleteRecommendation(req.Location); err != nil {
		log.Errorf("Failed to delete location[%v]: %v", req.Location, err)
		return err
	}

	return nil
}

func ListRecommendation() (response.ListRecommendationResponse, error) {
	var resp response.ListRecommendationResponse
	recommendatons, err := model.ListRecommendation()
	if err != nil {
		log.Errorf("Failed to list recommendations: %v", err)
		return resp, err
	}

	resp.Tickets = make([]response.Ticket, 0)
	for _, recommendation := range recommendatons {
		ticket, err := GetTicketByID(recommendation.TicketID)
		if err != nil {
			log.Errorf("Failed to get ticket by ID[%v]: %v", recommendation.TicketID, err)
			return resp, err
		}

		resTicket := response.Ticket{}
		resTicket.Load(ticket)
		resp.Tickets = append(resp.Tickets, resTicket)
	}

	return resp, nil
}
