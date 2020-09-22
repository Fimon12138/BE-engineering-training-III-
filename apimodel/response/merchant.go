package response

import (
	model2 "tickethub_service/model"
	"tickethub_service/util"
)

type GetMerchantResponse struct {
	ID            string `json:"id"`
	PayID         string `json:"payId"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Telephone     string `json:"telephone"`
	Email         string `json:"email"`
	Description   string `json:"description"`
	Qualified     string `json:"qualified"`
	QualifiedTime string `json:"qualifiedTime"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

type CreateMerchantResponse struct {
	ID          string `json:"id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
}

func (g *GetMerchantResponse) Load(merchant model2.Merchant) {
	g.ID = merchant.ID
	g.PayID = merchant.PayId
	g.Nickname = merchant.Nickname
	g.Avatar = merchant.Avatar
	g.Telephone = merchant.Telephone
	g.Email = merchant.Email
	g.Description = merchant.Description
	g.Qualified = merchant.Qualified
	g.QualifiedTime = util.GetDisplayTime(merchant.QualifiedTime)
	g.CreateTime = util.GetDisplayTime(merchant.CreateTime)
	g.UpdateTime = util.GetDisplayTime(merchant.UpdateTime)
}

func (c *CreateMerchantResponse) Load(merchant model2.Merchant) {
	c.ID = merchant.ID
	c.Nickname = merchant.Nickname
	c.Avatar = merchant.Avatar
	c.Telephone = merchant.Telephone
	c.Email = merchant.Email
	c.Description = merchant.Description
	c.CreateTime = util.GetDisplayTime(merchant.CreateTime)
}
