package response

import (
	model2 "tickethub_service/model"
	"tickethub_service/util"
)

type GetAccountResponse struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Type       string `json:"type"`
	RelatedID  string `json:"relatedId"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"update_time"`
}

type CreateAccountResponse struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Type       string `json:"type"`
	RelatedID  string `json:"relatedId"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"update_time"`
}

type LogIn struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (g *GetAccountResponse) Load(account model2.Account) {
	g.Name = account.Name
	g.Password = account.Password
	g.Type = account.Type
	g.RelatedID = account.RelatedID
	g.CreateTime = util.GetDisplayTime(account.CreateTime)
	g.UpdateTime = util.GetDisplayTime(account.UpdateTime)
}

func (c *CreateAccountResponse) Load(account model2.Account) {
	c.Name = account.Name
	c.Password = account.Password
	c.Type = account.Type
	c.RelatedID = account.RelatedID
	c.CreateTime = util.GetDisplayTime(account.CreateTime)
	c.CreateTime = util.GetDisplayTime(account.UpdateTime)
}
