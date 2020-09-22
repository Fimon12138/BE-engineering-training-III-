package request

type GetZjpayRequest struct {
	ID string `json:"id"`
}

type UpdateZjpayRequest struct {
	ID    string  `json:"id"`
	Money float32 `json:"money"`
}

type DeleteZjpayRequest struct {
	ID string `json:"id"`
}

type ChargeMoneyRequest struct {
	ID          string  `json:"id"`
	ChargeMoney float32 `json:"chargeMoney"`
}
