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

func CreateFavorite(req request.CreateFavoriteRequest) (response.CreateFavoriteResponse, error) {
	var resp response.CreateFavoriteResponse

	ticket, err := GetTicketByID(req.TicketID)
	if err != nil {
		log.Errorf("Failed to get ticket by ID[%v]: %v", req.TicketID, err)
		return resp, err
	}

	newFavorite := model.Favorite{
		ID:         util.NewUUIDString(enum.TABLENAME_FAVORITE),
		TicketID:   req.TicketID,
		UserID:     req.UserID,
		CreateTime: time.Now(),
	}

	favorite, err := model.CreateFavorite(newFavorite)
	if err != nil {
		log.Errorf("Failed to create favorite[%v]: %v", newFavorite, err)
		return resp, err
	}

	ticket.SubscribeCount++
	ticket.UpdateTime = time.Now()
	err = model.UpdateTicket(ticket)
	if err != nil {
		log.Errorf("Failed to update ticket[%v], err", ticket, err)
		return resp, err
	}

	resp.ID = favorite.ID
	resp.UserID = favorite.UserID
	resp.TicketID = favorite.TicketID
	return resp, err
}

func ListFavorite(req request.ListFavoriteRequest) (response.ListFavoriteResponse, error) {
	var resp response.ListFavoriteResponse

	filter := model.Favorite{
		UserID: req.UserID,
	}

	orderf := model.OrderFilter{
		Field:     req.OrderBy,
		Direction: req.Order,
	}

	page := model.Pagination{
		Offset: (req.PageNo - 1) * req.PageSize,
		Size:   req.PageSize,
	}

	favorites, totalCount, err := model.ListFavorite(filter, page, orderf)
	if err != nil {
		log.Errorf("Failed to list favorite by req[%v]: %v", req, err)
		return resp, nil
	}

	resp.Result = make([]response.Ticket, 0)
	for _, favorite := range favorites {
		ticket, err := GetTicketByID(favorite.TicketID)
		if err != nil {
			log.Errorf("Failed to get ticket by ID[%v]: %v", favorite.TicketID, err)
			return resp, err
		}

		newTicket := response.Ticket{}
		newTicket.Load(ticket)
		newTicket.IsSubscribed = enum.YES
		resp.Result = append(resp.Result, newTicket)
	}

	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.TotalCount = totalCount

	return resp, nil
}

func DeleteFavorite(req request.DeleteFavoriteRequest) error {
	ticket, err := GetTicketByID(req.TicketID)
	if err != nil {
		log.Errorf("Failed to get ticket by ID[%v]: %v", req.TicketID, err)
		return err
	}
	err = model.DeleteFavorite(model.Favorite{UserID: req.UserID,
												TicketID: req.TicketID})
	if err != nil {
		log.Errorf("Failed to delete favorite by req[%v]: %v", req, err)
		return err
	}
	ticket.SubscribeCount--
	ticket.UpdateTime = time.Now()
	err = model.UpdateTicket(ticket)
	if err != nil {
		log.Errorf("Failed to update ticket[%v]: %v", ticket, err)
		return err
	}
	return nil
}
