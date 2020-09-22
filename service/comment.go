package service

import (
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/response"
	"tickethub_service/model"
	"tickethub_service/util"
	"tickethub_service/util/enum"
	"tickethub_service/util/log"
	"time"
)

func CreateComment(req request.CreateCommentRequest) (response.CreateCommentResponse, error) {
	var resp response.CreateCommentResponse

	ticket, err := GetTicketByID(req.TicketID)
	if err != nil {
		log.Errorf("Failed to get Ticket by ID[%v]: %v", req.TicketID, err)
		return resp, err
	}

	user, err := GetUserByID(req.UserID)
	if err != nil {
		log.Errorf("Failed to get user by ID[%v]: %v", req.UserID, err)
		return resp, err
	}

	newComment := model.Comment{
		ID:         util.NewUUIDString(enum.TABLENAME_COMMENT),
		TicketID:   req.TicketID,
		TicketName: ticket.Name,
		UserID:     req.UserID,
		UserName:   user.Nickname,
		Content:    req.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	comment, err := model.CreateComment(newComment)
	if err != nil {
		log.Errorf("Failed to create comment[%v]: %v", newComment, err)
		return resp, err
	}

	resp.Load(comment)
	return resp, nil
}

func ListComment(req request.ListCommentRequest) (response.ListCommentResponse, error) {
	var resp response.ListCommentResponse

	orderf := model.OrderFilter{
		Field:     enum.ORDERBY_COMMON_UPDATETIME,
		Direction: req.Order,
	}

	page := model.Pagination{
		Offset: (req.PageNo - 1) * req.PageSize,
		Size:   req.PageSize,
	}

	filter := model.Comment{
		UserID:   req.UserID,
		TicketID: req.TicketID,
	}

	comments, totalCount, err := model.ListComment(filter, page, orderf)
	if err != nil {
		log.Errorf("Failed to ListComment by req[%v]: %v", req, err)
		return resp, err
	}

	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Result = make([]response.Comment, 0)
	for _, comment := range comments {
		newComment := response.Comment{}
		newComment.Load(comment)
		resp.Result = append(resp.Result, newComment)
	}
	resp.TotalCount = totalCount

	return resp, nil
}

func UpdateComment(req request.UpdateCommentRequest) error {
	comment, err := GetCommentByID(req.ID)
	if err != nil {
		log.Errorf("Failed to get comment by ID[%v]: %v", req.ID, err)
		return err
	}

	comment.Content = req.Content
	comment.UpdateTime = time.Now()
	err = model.UpdateComment(comment)

	if err != nil {
		log.Errorf("Failed to update comment[%v]: %v", comment, err)
		return err
	}
	return nil
}

func DeleteComment(req request.DeleteCommentRequest) error {
	err := model.DeleteComment(req.ID)
	if err != nil {
		log.Errorf("Failed to delete comment by ID[%v]: %v", req.ID, err)
		return err
	}

	return nil
}

func GetCommentByID(ID string) (model.Comment, error) {
	return model.GetComment(ID)
}
