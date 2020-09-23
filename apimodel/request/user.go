package request

type GetUserRequest struct {
	ID string `json:"id"`
}

type CreateUserRequest struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Description string `json:"description"`
}

type UpdateUserRequest struct {
	ID          string `json:"id"`
	PayID       string `json:"payId"`
	NickName    string `json:"nickName"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Description string `json:"Description"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}
