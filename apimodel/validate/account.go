package validate

import (
	"fmt"
	request2 "tickethub_service/apimodel/request"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/errors"
)

func CheckCreateAccount(req request2.CreateAccountRequest) error {
	if !util.ContainsString(enum.GetAccountType(), req.Type) {
		msg := fmt.Sprintf("The content of accountType[%v] wrong", req.Type)
		return errors.InvalidRequestError(msg)
	}
	return nil
}
