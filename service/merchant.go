package service

import (
	request2 "tickethub_service/apimodel/request"
	response2 "tickethub_service/apimodel/response"
	model2 "tickethub_service/model"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/log"
	"time"
)

func GetMerchant(req request2.GetMerchantRequest) (response2.GetMerchantResponse, error) {
	var resp response2.GetMerchantResponse

	merchant, err := model2.GetMerchant(req.ID)
	if err != nil {
		log.Errorf("Failed to GetMerchant by id [%v]: %v", req.ID, err)
		return resp, err
	}

	resp.Load(merchant)
	return resp, nil
}

func CreateMerchant(req request2.CreateMerchantRequest) (response2.CreateMerchantResponse, error) {
	var resp response2.CreateMerchantResponse

	newMerchant := model2.Merchant{
		ID:            util.NewUUIDString(enum.TABLENAME_MERCHANT),
		Nickname:      req.Nickname,
		Avatar:        req.Avatar,
		Telephone:     req.Telephone,
		Email:         req.Email,
		Qualified:     enum.NO,
		QualifiedTime: time.Now(),
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	merchant, err := model2.CreateMerchant(newMerchant)
	if err != nil {
		log.Errorf("Failed to createMerchant in database: %v", err)
		return resp, err
	}

	resp.Load(merchant)
	return resp, nil
}

func UpdateMerchant(req request2.UpdateMerchantRequest) error {
	merchant, err := GetMerchantByID(req.ID)
	if err != nil {
		log.Errorf("Failed to GetMerchant by id [%v] when update: %v", req.ID, err)
		return err
	}

	merchant.PayId = req.PayID
	merchant.UpdateTime = time.Now()
	merchant.Nickname = req.Nickname
	merchant.Avatar = req.Avatar
	merchant.Telephone = req.Telephone
	merchant.Email = req.Email
	merchant.Description = req.Description
	merchant.Qualified = req.Qualified

	if err := model2.UpdateMerchant(merchant); err != nil {
		log.Errorf("Failed to updateMerchant by req[%v]: %v", req, err)
		return err
	}

	return nil
}

func DeleteMerchant(req request2.DeleteMerchantRequest) error {
	err := model2.DeleteMerchant(req.ID)
	if err != nil {
		log.Errorf("Failed to deleteMerchant by req[%v]: %v", req, err)
		return err
	}

	return nil
}

func GetMerchantByID(ID string) (model2.Merchant, error) {
	return model2.GetMerchant(ID)
}
