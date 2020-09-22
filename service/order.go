package service

import (
	"fmt"
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/response"
	"tickethub_service/model"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
	"time"
)

func GetOrder(req request.GetOrderRequest) (response.GetOrderResponse, error) {
	var resp response.GetOrderResponse
	order, err := model.GetOrder(req.ID)
	if err != nil {
		log.Errorf("Failed t get order by ID[%v]: %v", req.ID, err)
		return response.GetOrderResponse{}, err
	}

	resp.Load(order)
	return resp, nil
}

func CreateOrder(req request.CreateOrderRequest) (response.CreateOrderResponse, error) {
	var resp response.CreateOrderResponse

	ticket, err := GetTicketByID(req.TicketID)
	if err != nil {
		log.Errorf("Failed to get Ticket by ID[%v]: %v", req.TicketID, err)
		return resp, err
	}

	user, err := GetUserByID(req.UserID)
	if err != nil {
		log.Errorf("Failed to get user by ID[%v]: %v", req.UserID, err)
		return resp, err
	}

	newOrder := model.Order{
		ID:         util.NewUUIDString("order"),
		TicketID:   req.TicketID,
		TicketName: ticket.Name,
		UserID:     req.UserID,
		UserName:   user.Nickname,
		Status:     enum.ORDER_STATUS_PROCESSING,
		Price:      req.Price,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	order, err := model.CreateOrder(newOrder)
	if err != nil {
		log.Errorf("Failed to create order[%v]: %v", newOrder, err)
		return resp, err
	}
	resp.Load(order)

	return resp, nil
}

func UpdateOrder(req request.UpdateOrderRequest) error {
	order, err := GetOrderByID(req.ID)
	if err != nil {
		log.Errorf("Failed to get order by ID[%v]: %v", req.ID, err)
		return nil
	}

	order.Status = req.Status
	order.UpdateTime = time.Now()

	err = model.UpdateOrder(order)
	if err != nil {
		log.Errorf("Failed to update order[%v]: %v", order, err)
		return err
	}

	return nil
}

func ListOrder(req request.ListOrderRequest) (response.ListOrderResponse, error) {
	var resp response.ListOrderResponse

	orderf := model.OrderFilter{
		Field:     enum.ORDERBY_COMMON_UPDATETIME,
		Direction: req.Order,
	}

	page := model.Pagination{
		Offset: (req.PageNo - 1) * req.PageSize,
		Size:   req.PageSize,
	}

	filter := model.Order{
		UserID: req.UserID,
		Status: req.Status,
	}

	orders, totalCount, err := model.ListOrders(filter, page, orderf)
	if err != nil {
		log.Errorf("Failed tp list orders by req[%v]: %v", req, err)
		return resp, err
	}
	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Result = make([]response.Order, 0)
	for _, order := range orders {
		newOrder := response.Order{}
		newOrder.Load(order)
		resp.Result = append(resp.Result, newOrder)
	}
	resp.TotalCount = totalCount

	return resp, nil
}

func GetOrderByID(ID string) (model.Order, error) {
	return model.GetOrder(ID)
}

func PayForOrder(req request.PayForOrderRequest) error {
	order, err := GetOrderByID(req.OrderID)
	if err != nil {
		log.Errorf("Failed to get order by ID[%v]: %v", req.OrderID, err)
		return err
	}

	ticket, err := GetTicketByID(order.TicketID)
	if err != nil {
		log.Errorf("Failed to get ticket by ID[%v]: %v", order.TicketID, err)
		return err
	}

	zjpay, err := GetZjpayByID(req.ZjpayID)
	if err != nil {
		log.Errorf("Failed to get zjpay by ID[%v]: %v", req.ZjpayID, err)
		return err
	}

	if zjpay.Money < order.Price {
		msg := fmt.Sprintf("Don't have enough money zjpay[%v] order[%v]", zjpay.Money, order.Price)
		log.Errorf(msg)
		return errors.InternalError(msg)
	}

	order.Status = enum.ORDER_STATUS_FINISHED
	order.UpdateTime = time.Now()

	ticket.Count -= 1
	ticket.UpdateTime = time.Now()

	zjpay.Money -= order.Price
	zjpay.UpdateTime = time.Now()

	err = model.UpdateOrder(order)
	if err != nil {
		log.Errorf("Failed to update order by req[%v]: %v", order, err)
		return err
	}

	err = model.UpdateTicket(ticket)
	if err != nil {
		log.Errorf("Failed to update ticket by req[%v]: %v", order, err)
		return err
	}

	err = model.UpdateZjpay(zjpay)
	if err != nil {
		log.Errorf("Failed to update zjpay by req[%v]: %v", zjpay, err)
		return err
	}
	return nil
}

func ListFinishedTicket(req request.ListFinishedTicketRequest) (response.ListFinishedTicketResponse, error) {
	var resp response.ListFinishedTicketResponse

	filter := model.Order{
		UserID: req.UserID,
		Status: enum.ORDER_STATUS_FINISHED,
	}

	orderf := model.OrderFilter{
		Field:     req.OrderBy,
		Direction: req.Order,
	}

	page := model.Pagination{
		Offset: (req.PageNo - 1) * req.PageSize,
	}

	orders, totalCount, err := model.ListOrders(filter, page, orderf)
	if err != nil {
		log.Errorf("Failed to get orders by req[%v]: %v", req, err)
		return resp, err
	}

	resp.Result = make([]response.Ticket, 0)
	for _, order := range orders {
		ticket, err := GetTicketByID(order.TicketID)
		if err != nil {
			log.Errorf("Failed to get ticket by ID[%v]: %v", order.TicketID, err)
			return resp, err
		}
		respTicket := response.Ticket{}
		respTicket.Load(ticket)
		resp.Result = append(resp.Result, respTicket)
	}

	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.TotalCount = totalCount
	return resp, nil

}
