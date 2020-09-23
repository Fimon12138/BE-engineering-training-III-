package response

import (
	"tickethub_service/model"
	"tickethub_service/util"
)

type Comment struct {
	ID         string `json:"id"`
	TicketID   string `json:"ticketId"`
	TicketName string `json:"ticketName"`
	UserID     string `json:"userId"`
	UserName   string `json:"userName"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type GetCommentResponse struct {
	Comment
}

type CreateCommentResponse struct {
	Comment
}

type ListCommentResponse struct {
	PageNo     int       `json:"pageNo"`
	PageSize   int       `json:"pageSize"`
	Result     []Comment `json:"result"`
	TotalCount int       `json:"totalCount"`
}

func (c *Comment) Load(comment model.Comment) {
	c.ID = comment.ID
	c.TicketID = comment.TicketID
	c.TicketName = comment.TicketName
	c.UserID = comment.UserID
	c.UserName = comment.UserName
	c.Content = comment.Content
	c.CreateTime = util.GetDisplayTime(comment.CreateTime)
	c.UpdateTime = util.GetDisplayTime(comment.UpdateTime)
}
