package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type User struct {
	ID          string `gorm:"id"`
	PayID       string `gorm:"pay_id"`
	Nickname    string
	Avatar      string
	Telephone   string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
}

func GetUser(ID string) (User, error) {
	var user User

	err := config.DB.Where(User{ID: ID}).Find(&user).Error
	return user, err
}

func CreateUser(newUser User) (User, error) {
	if err := config.DB.Table(enum.TABLENAME_USER).Create(newUser).Error; err != nil {
		return User{}, err
	}
	return newUser, nil
}

func UpdateUser(newUser User) error {
	return config.DB.Table(enum.TABLENAME_USER).Where(&User{ID: newUser.ID}).Update(newUser).Error
}

func DeleteUser(ID string) error {
	return config.DB.Table(enum.TABLENAME_USER).Delete(&User{ID: ID}).Error
}
