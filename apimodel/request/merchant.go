package request

type GetMerchantRequest struct {
	ID string `json:"id"`
}

type CreateMerchantRequest struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Description string `json:"Description"`
}

type UpdateMerchantRequest struct {
	ID          string `json:"id" binding:"required"`
	PayID       string `json:"payId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Description string `json:"Description"`
	Qualified   string `json:"qualified"`
}

type DeleteMerchantRequest struct {
	ID string `json:"id"`
}
