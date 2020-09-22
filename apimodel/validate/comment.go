package validate

import (
	"fmt"
	"tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
)

func CheckListCommentRequest(req *request.ListCommentRequest) error {
	if !CheckPages(req.PageSize, req.PageNo) {
		msg := fmt.Sprintf("The content of PageNo[%v] PageSize[%v] wrong", req.PageNo, req.PageSize)
		return errors.InvalidRequestError(msg)
	}

	if !util.ContainsStringAllowZero(enum.GetOrderValues(), req.Order) {
		msg := fmt.Sprintf("The content of Order[%v] wrong", req.Order)
		return errors.InvalidRequestError(msg)
	}

	if req.Order == "" {
		req.Order = enum.ORDER_DESC
	}

	req.OrderBy = enum.ORDERBY_COMMON_UPDATETIME
	return nil
}
