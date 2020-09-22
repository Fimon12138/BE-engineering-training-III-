package validate

import (
	"fmt"
	request2 "tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
)

func CheckCreateMerchant(req request2.CreateMerchantRequest) error {
	isMatch, err := util.CheckEmaliAddress(req.Email)
	if err != nil {
		msg := fmt.Sprintf("check email failed:%v", err)
		log.Errorf(msg)
		return errors.InternalError(msg)
	}
	if !util.CheckTelephone(req.Telephone) || !isMatch {
		msg := fmt.Sprintf("content of telephone[%v] or email[%v] wrong", req.Telephone, req.Email)
		log.Errorf(msg)
		return errors.InvalidRequestError(msg)
	}
	return nil
}

func CheckUpdateMerchant(req request2.UpdateMerchantRequest) error {
	isMatch, err := util.CheckEmaliAddress(req.Email)
	if err != nil {
		msg := fmt.Sprintf("check email failed:%v", err)
		log.Errorf(msg)
		return errors.InternalError(msg)
	}
	if !util.CheckTelephone(req.Telephone) || !isMatch || !util.ContainsStringAllowZero(enum.GetBoolType(), req.Qualified) {
		msg := fmt.Sprintf("content of telephone[%v] or email[%v] or quilified[%v] wrong", req.Telephone, req.Email, req.Qualified)
		log.Errorf(msg)
		return errors.InvalidRequestError(msg)
	}

	return nil
}
