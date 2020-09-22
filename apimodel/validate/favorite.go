package validate

import (
	"fmt"
	"tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
)

func CheckListFavorite(req *request.ListFavoriteRequest) error {
	if !CheckPages(req.PageNo, req.PageSize) {
		msg := fmt.Sprintf("The content of pageNo[%v] or pageSize[%v] wrong", req.PageNo, req.PageSize)
		return errors.InvalidRequestError(msg)
	}

	if !util.ContainsStringAllowZero(enum.GetOrderValues(), req.Order) {
		msg := fmt.Sprintf("The content of order[%v] wrong", req.Order)
		return errors.InvalidRequestError(msg)
	}

	if req.Order == "" {
		req.Order = enum.ORDER_DESC
	}

	req.OrderBy = enum.ORDERBY_COMMON_CREATETIME

	return nil
}
