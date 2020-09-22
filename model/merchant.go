package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Merchant struct {
	ID            string `gorm:"id"`
	PayId         string `gorm:"pay_id"`
	Nickname      string
	Avatar        string
	Telephone     string
	Email         string
	Description   string
	Qualified     string
	QualifiedTime time.Time
	CreateTime    time.Time
	UpdateTime    time.Time
}

func GetMerchant(ID string) (Merchant, error) {
	var merchant Merchant
	err := config.DB.Where(Merchant{ID: ID}).Find(&merchant).Error
	return merchant, err
}

func CreateMerchant(newMerchant Merchant) (Merchant, error) {
	err := config.DB.Table(enum.TABLENAME_MERCHANT).Create(newMerchant).Error
	if err != nil {
		return Merchant{}, err
	}
	return newMerchant, nil
}

func UpdateMerchant(newMerchant Merchant) error {
	return config.DB.Table(enum.TABLENAME_MERCHANT).Where(&Merchant{ID: newMerchant.ID}).Update(newMerchant).Error
}

func DeleteMerchant(ID string) error {
	return config.DB.Table(enum.TABLENAME_MERCHANT).Delete(&Merchant{ID: ID}).Error
}
