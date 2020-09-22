package service

import (
	"tickethub_service/apimodel/request"
	"tickethub_service/model"
	"tickethub_service/util"
	"tickethub_service/util/log"
	"time"
)

func GetZjpay(req request.GetZjpayRequest) (model.Zjpay, error) {

	resp, err := model.GetZjpay(req.ID)
	if err != nil {
		log.Errorf("Failed to get zjpay by id[%v]: %v", req.ID, err)
		return model.Zjpay{}, err
	}

	return resp, nil
}

func CreateZjPay() (model.Zjpay, error) {
	zjpay := model.Zjpay{
		ID:         util.NewUUIDString("zjpay"),
		Money:      0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	newZjpay, err := model.CreateZjpay(zjpay)
	if err != nil {
		log.Errorf("Failed to CreateZjpay by zjpay[%v]: %v", zjpay, err)
		return model.Zjpay{}, err
	}

	return newZjpay, err
}

func UpdateZjPay(req request.UpdateZjpayRequest) error {
	zjpay, err := GetZjpayByID(req.ID)
	if err != nil {
		log.Errorf("Failed to get zjpay by id[%v]: %v", req.ID, err)
		return err
	}

	zjpay.Money = req.Money
	zjpay.UpdateTime = time.Now()

	err = model.UpdateZjpay(zjpay)
	if err != nil {
		log.Errorf("Failed to UpdateZjpay by zjpay[%v]: %v", req, err)
		return err
	}

	return nil
}

func DeleteZjpay(req request.DeleteZjpayRequest) error {
	err := model.DeleteZjpay(req.ID)
	if err != nil {
		log.Errorf("Failed to delete zjpay by req[%v]: %v", req, err)
		return err
	}

	return nil
}

func ChargeMoney(req request.ChargeMoneyRequest) error {
	zjpay, err := GetZjpayByID(req.ID)
	if err != nil {
		log.Errorf("Failed to get zjpay by ID[%v]: [%v]", req.ID, err)
		return err
	}
	zjpay.Money += req.ChargeMoney
	zjpay.UpdateTime = time.Now()
	err = model.UpdateZjpay(zjpay)
	if err != nil {
		log.Errorf("Failed to uodate zjpay by zjpay[%v]: [%v]", zjpay, err)
		return err
	}

	return nil
}

func GetZjpayByID(ID string) (model.Zjpay, error) {
	return model.GetZjpay(ID)
}
