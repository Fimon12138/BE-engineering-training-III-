package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Zjpay struct {
	ID         string `gorm:"id"`
	Money      float32
	CreateTime time.Time
	UpdateTime time.Time
}

func GetZjpay(ID string) (Zjpay, error) {
	var zjpay Zjpay
	err := config.DB.Table(enum.TABLENAME_ZJPAY).Where(&Zjpay{ID: ID}).Find(&zjpay).Error
	return zjpay, err
}
func CreateZjpay(newZjpay Zjpay) (Zjpay, error) {
	if err := config.DB.Table(enum.TABLENAME_ZJPAY).Create(newZjpay).Error; err != nil {
		return Zjpay{}, err
	}

	return newZjpay, nil
}

func UpdateZjpay(newZjpay Zjpay) error {
	return config.DB.Table(enum.TABLENAME_ZJPAY).Where(&Zjpay{ID: newZjpay.ID}).Update(newZjpay).Error
}

func DeleteZjpay(ID string) error {
	return config.DB.Table(enum.TABLENAME_ZJPAY).Delete(&Zjpay{ID: ID}).Error
}
