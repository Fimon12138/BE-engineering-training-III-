package response

import (
	model2 "tickethub_service/model"
	"tickethub_service/util"
)

type GetUserResponse struct {
	ID          string `json:"id"`
	PayId       string `json:"payId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Description string `json:"description"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type CreateUserResponse struct {
	ID          string `json:"id"`
	PayId       string `json:"payId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Description string `json:"description"`
	CreateTime  string `json:"create_time"`
}

func (g *GetUserResponse) Load(user model2.User) {
	g.ID = user.ID
	g.PayId = user.PayID
	g.Avatar = user.Avatar
	g.Telephone = user.Telephone
	g.Nickname = user.Nickname
	g.Description = user.Description
	g.CreateTime = util.GetDisplayTime(user.CreateTime)
	g.UpdateTime = util.GetDisplayTime(user.UpdateTime)
}

func (c *CreateUserResponse) Load(user model2.User) {
	c.ID = user.ID
	c.PayId = user.PayID
	c.Avatar = user.Avatar
	c.Telephone = user.Telephone
	c.Nickname = user.Nickname
	c.Description = user.Description
	c.CreateTime = util.GetDisplayTime(user.CreateTime)
}
