package validate

import (
	"fmt"
	"tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
)

func CheckListTicket(req *request.ListTicketRequest) error {
	if !CheckPages(req.PageNo, req.PageSize) {
		msg := fmt.Sprintf("The content of PageNo[%v] PageSize[%v] wrong", req.PageNo, req.PageSize)
		return errors.InvalidRequestError(msg)
	}

	if !util.ContainsStringAllowZero(enum.GetTicketTypeValues(), req.Type) ||
		!util.ContainsStringAllowZero(enum.GetTicketOrderByValues(), req.OrderBy) ||
		!util.ContainsStringAllowZero(enum.GetOrderValues(), req.Order) ||
		!util.ContainsStringAllowZero(enum.GetBoolType(), req.OutOfDate) {
		msg := fmt.Sprintf("The content of Type[%v] or "+
			"orderby[%v] or order[%v] or outOfDate[%v] wrong", req.Type, req.OrderBy, req.Order, req.OutOfDate)
		return errors.InvalidRequestError(msg)
	}
	if req.Order == "" {
		req.Order = enum.ORDER_DESC
	}

	if req.OrderBy == "" {
		req.OrderBy = enum.ORDERBY_COMMON_UPDATETIME
	}
	return nil
}
