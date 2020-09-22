package service

import (
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/response"
	"tickethub_service/model"
	"tickethub_service/util/log"
	"time"
)

func GetAccount(req request.GetAccountRequest) (response.GetAccountResponse, error) {
	var resp response.GetAccountResponse
	account, err := model.GetAccount(req.AccountName)
	if err != nil {
		log.Errorf("Failed to get Account by AssetName[%v]: %v", req.AccountName, err)
		return resp, err
	}

	resp.Load(account)
	return resp, nil
}

func CreateAccount(req request.CreateAccountRequest) (response.CreateAccountResponse, error) {
	var resp response.CreateAccountResponse

	newAccount := model.Account{
		Name:       req.Name,
		Password:   req.Password,
		Type:       req.Type,
		RelatedID:  req.RelatedID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	account, err := model.CreateAccount(newAccount)
	if err != nil {
		log.Errorf("Failed to create account by req[%v]: %v", req, err)
		return resp, err
	}
	resp.Load(account)
	return resp, nil
}

func UpdateAccount(req request.UpdateAccountRequest) error {
	account, err := GetAccountByName(req.Name)
	if err != nil {
		log.Errorf("Failed to updateAccount by req[%v]: %v", req, err)
		return err
	}

	account.Password = req.Password
	account.UpdateTime = time.Now()

	err = model.UpdateAccount(account)
	if err != nil {
		log.Errorf("Failed to updateAccount by req[%v]: %v", req, err)
		return err
	}
	return nil
}

func DeleteAccount(req request.DeleteAccountRequest) error {
	err := model.DeleteAccount(req.Name)
	if err != nil {
		log.Errorf("Failed to delete account by name[%v]: %v", req.Name, err)
		return err
	}
	return nil
}

func LogIn(req request.LogIn) response.LogIn {
	var resp response.LogIn
	account, err := model.GetAccount(req.Name)
	if err != nil || account.Password != req.Password {
		return resp
	}
	//TODO add token
	resp = response.LogIn{}
	resp.ID = account.RelatedID
	return resp
}

func GetAccountByName(name string) (model.Account, error) {
	return model.GetAccount(name)
}
