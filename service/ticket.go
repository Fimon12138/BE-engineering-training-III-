package service

import (
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/response"
	"tickethub_service/model"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/log"
	"time"
)

func GetTicket(req request.GetTicketRequest) (response.GetTicketResponse, error) {
	var resp response.GetTicketResponse

	ticket, err := model.GetTicket(req.ID)
	if err != nil {
		log.Errorf("Failed to get ticket by ID[%v]: %v", req.ID, err)
		return response.GetTicketResponse{}, err
	}

	if req.UserID != "" {
		filter := model.Favorite{
			UserID: req.UserID,
			TicketID: ticket.ID,
		}
		isFavorite, err := model.MatchFavorite(filter)
		if err != nil {
			log.Errorf("Failed to matchFavorite userID[%v] ticketID[%v]", req.UserID, ticket.ID)
			return resp, err
		}

		if isFavorite {
			resp.IsSubscribed = enum.YES
		}	else {
			resp.IsSubscribed = enum.NO
		}
	}

	resp.Load(ticket)
	return resp, nil
}

func ListTicket(req request.ListTicketRequest) (response.ListTicketResponse, error) {
	var resp response.ListTicketResponse

	orderf := model.OrderFilter{
		Field:     req.OrderBy,
		Direction: req.Order,
	}

	page := model.Pagination{
		Offset: (req.PageNo - 1) * req.PageSize,
		Size:   req.PageSize,
	}

	filter := model.Ticket{
		Type: req.Type,
	}

	tickets, totalCount, err := model.ListTicket(filter, page, orderf, req.NameFilter)
	if err != nil {
		log.Errorf("Failed to List ticket by req[%v]: %v", req, err)
		return resp, err
	}

	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Result = make([]response.Ticket, 0)
	for _, ticket := range tickets {
		newTicket := response.Ticket{}
		newTicket.Load(ticket)
		resp.Result = append(resp.Result, newTicket)
	}
	resp.TotalCount = totalCount

	return resp, nil
}

func CreateTicket(req request.CreateTicketRequest) (response.CreateTicketResponse, error) {
	var resp response.CreateTicketResponse
	startTime, err := util.GetTimeFromString(req.StartTime)
	if err != nil {
		log.Errorf("Failed to get time form string[%v]: %v", req.StartTime, err)
		return resp, err
	}

	newTicket := model.Ticket{
		ID:          util.NewUUIDString(enum.TABLENAME_TICKET),
		Name:        req.Name,
		MerchantID:  req.MerchantID,
		ImageRow:    req.ImageRow,
		ImageColumn: req.ImageColumn,
		ImageDetail: req.ImageDetail,
		Type:        req.Type,
		Price:       req.Price,
		StartTime:   startTime,
		Location:    req.Location,
		Count:       req.Count,
		Description: req.Description,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	ticket, err := model.CreateTicket(newTicket)
	if err != nil {
		log.Errorf("Failed to create ticket[%v] in db:%v", newTicket, err)
		return resp, err
	}
	resp.Load(ticket)
	return resp, nil
}

func UpdateTicket(req request.UpdateTicketRequest) error {
	ticket, err := GetTicketByID(req.ID)
	if err != nil {
		log.Errorf("Failed to get Ticket by ID[%v]: %v", req.ID, err)
		return err
	}

	newTicket := ticket
	newTicket.Name = req.Name
	newTicket.Price = req.Price
	newTicket.Count = req.Count
	if req.SubscribeIncrease {
		newTicket.SubscribeCount++
	}

	err = model.UpdateTicket(newTicket)
	if err != nil {
		log.Errorf("Failed to update ticket[%v]: %v", ticket, err)
		return err
	}

	return nil
}

func DeleteTicket(req request.DeleteTicketRequest) error {
	err := model.DeleteTicket(req.ID)
	if err != nil {
		log.Errorf("Failed to delete ticket by ID[%v]: %v", req.ID, err)
		return err
	}
	return nil
}

func GetTicketByID(ID string) (model.Ticket, error) {
	return model.GetTicket(ID)
}
