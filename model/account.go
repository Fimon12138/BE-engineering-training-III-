package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Account struct {
	Name       string
	Password   string
	Type       string
	RelatedID  string `gorm:"related_id"`
	CreateTime time.Time
	UpdateTime time.Time
}

func GetAccount(accountName string) (Account, error) {
	var account Account
	database := config.DB.Where(&Account{Name: accountName})

	database.Find(&account)
	if database.Error != nil {
		return Account{}, database.Error
	}

	return account, nil
}

func CreateAccount(newAccount Account) (Account, error) {
	err := config.DB.Table(enum.TABLENAME_ACCOUNT).Create(newAccount).Error
	if err.Error != nil {
		return Account{}, err
	}
	return newAccount, nil
}

func UpdateAccount(newAccount Account) error {
	return config.DB.Where(&Account{Name: newAccount.Name}).Update(newAccount).Error
}

func DeleteAccount(accountName string) error {
	return config.DB.Table(enum.TABLENAME_ACCOUNT).Delete(&Account{Name: accountName}).Error
}
