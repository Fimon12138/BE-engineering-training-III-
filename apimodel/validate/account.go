package validate

import (
	"fmt"
	"tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
)

func CheckCreateAccount(req request.CreateAccountRequest) error {
	if !util.ContainsString(enum.GetAccountType(), req.Type) {
		msg := fmt.Sprintf("The content of accountType[%v] wrong", req.Type)
		return errors.InvalidRequestError(msg)
	}
	return nil
}

func CheckSignup(req request.SignUp) error{
	if !util.CheckTelephone(req.Telephone) {
		msg := fmt.Sprintf("The content og Telephone[%v] wrong", req.Telephone)
		return errors.InvalidRequestError(msg)
	}

	if !util.ContainsString(enum.GetAccountType(), req.Type) {
		msg := fmt.Sprintf("The content of Type[%v] wrong", req.Type)
		return errors.InvalidRequestError(msg)
	}

	return nil
}