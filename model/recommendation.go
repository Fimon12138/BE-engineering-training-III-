package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Recommendation struct {
	TicketID   string `gorm:"ticket_id"`
	Location   int
	CreateTime time.Time
	UpdateTime time.Time
}

func GetRecommendationByLocation(location int) (Recommendation, error) {
	var recommendation Recommendation

	err := config.DB.Where(&Recommendation{Location: location}).Find(&recommendation).Error
	if err != nil {
		return Recommendation{}, err
	}
	return recommendation, nil
}

func CreateRecommendation(recommendation Recommendation) error {
	return config.DB.Table(enum.TABLENAME_RECOMMENDATION).Create(recommendation).Error
}

func ListRecommendation() ([]Recommendation, error) {
	var recommendations []Recommendation

	err := config.DB.Table(enum.TABLENAME_RECOMMENDATION).
		Where(&Recommendation{}).Order("location" + " " + enum.ORDER_ASC).Find(&recommendations).Error

	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func DeleteRecommendation(location int) error {
	return config.DB.Table(enum.TABLENAME_RECOMMENDATION).Delete(Recommendation{Location: location}).Error
}
