package response

import (
	model2 "tickethub_service/model"
	"tickethub_service/util"
)

type GetZjpayResponse struct {
	ID         string  `json:"id"`
	Money      float32 `json:"money"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
}

type CreateZjpayResponse struct {
	ID         string `json:"id"`
	CreateTime string `json:"createTime"`
}

func (g *GetZjpayResponse) Load(zjpay model2.Zjpay) {
	g.ID = zjpay.ID
	g.Money = zjpay.Money
	g.CreateTime = util.GetDisplayTime(zjpay.CreateTime)
	g.UpdateTime = util.GetDisplayTime(zjpay.UpdateTime)
}

func (c *CreateZjpayResponse) Load(zjpay model2.Zjpay) {
	c.ID = zjpay.ID
	c.CreateTime = util.GetDisplayTime(zjpay.CreateTime)
}
