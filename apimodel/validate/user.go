package validate

import (
	"fmt"
	request2 "tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/errors"
)

func CheckCreateUser(req request2.CreateUserRequest) error {
	if !util.CheckTelephone(req.Telephone) {
		msg := fmt.Sprintf("Tellephone content[%v] wrong ", req.Telephone)
		return errors.InvalidRequestError(msg)
	}

	return nil
}

func CheckUpdateUser(req request2.UpdateUserRequest) error {
	if !util.CheckTelephone(req.Telephone) || req.Telephone == "" {
		msg := fmt.Sprintf("Tellephone content[%v] wrong ", req.Telephone)
		return errors.InvalidRequestError(msg)
	}

	return nil
}
