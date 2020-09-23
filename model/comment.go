package model

import (
	"tickethub_service/config"
	"tickethub_service/util/enum"
	"time"
)

type Comment struct {
	ID         string `gorm:"id"`
	TicketID   string `gorm:"ticket_id"`
	TicketName string
	UserID     string `gorm:"user_id"`
	UserName   string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetComment(ID string) (Comment, error) {
	var comment Comment

	err := config.DB.Table(enum.TABLENAME_COMMENT).Where(&Comment{ID: ID}).Find(&comment).Error

	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func CreateComment(newComment Comment) (Comment, error) {
	err := config.DB.Table(enum.TABLENAME_COMMENT).Create(newComment).Error
	if err != nil {
		return Comment{}, err
	}

	return newComment, err
}

func ListComment(filter Comment, p Pagination, orderf OrderFilter) ([]Comment, int, error) {
	var comments []Comment
	var totalCount int

	database := config.DB.Table(enum.TABLENAME_COMMENT).Where(&filter)
	database.Count(&totalCount)

	database = database.Order(orderf.Field + " " + orderf.Direction)
	database = database.Offset(p.Offset).Limit(p.Size)
	database.Find(&comments)

	if database.Error != nil {
		return nil, 0, database.Error
	}
	return comments, totalCount, nil
}

func UpdateComment(newComment Comment) error {
	return config.DB.Table(enum.TABLENAME_COMMENT).Where(&Comment{ID: newComment.ID}).Update(newComment).Error
}

func DeleteComment(ID string) error {
	return config.DB.Table(enum.TABLENAME_COMMENT).Delete(&Comment{ID: ID}).Error
}
