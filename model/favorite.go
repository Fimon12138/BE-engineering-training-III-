package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Favorite struct {
	ID         string `gorm:"id"`
	UserID     string `gorm:"user_id"`
	TicketID   string `gorm:"ticket_id"`
	CreateTime time.Time
}

func GetFavorite(ID string) (Favorite, error) {
	var favorete Favorite
	err := config.DB.Table(enum.TABLENAME_FAVORITE).Where(Favorite{ID: ID}).Find(&favorete).Error
	if err != nil {
		return favorete, err
	}

	return favorete, nil
}

func CreateFavorite(favo Favorite) (Favorite, error) {
	err := config.DB.Table(enum.TABLENAME_FAVORITE).Create(favo).Error
	if err != nil {
		return Favorite{}, err
	}
	return favo, nil
}

func ListFavorite(filter Favorite, page Pagination, orderf OrderFilter) ([]Favorite, int, error) {
	var favorites []Favorite
	var totalCount int

	database := config.DB.Table(enum.TABLENAME_FAVORITE).Where(&filter)
	database.Count(&totalCount)

	database = database.Order(orderf.Field + " " + orderf.Direction)
	database = database.Offset(page.Offset).Limit(page.Size)
	database.Find(&favorites)
	if database.Error != nil {
		return nil, 0, database.Error
	}

	return favorites, totalCount, nil
}

func DeleteFavorite(filter Favorite) error {
	return config.DB.Table(enum.TABLENAME_FAVORITE).Delete(&filter).Error
}

func MatchFavorite(filter Favorite) (bool, error) {
	var count int
	database := config.DB.Table(enum.TABLENAME_FAVORITE).Where(&filter)
	database.Count(&count)
	if database.Error != nil {
		return false, database.Error
	}
	if count <= 0 {
		return false, nil
	}
	return true, nil
}
