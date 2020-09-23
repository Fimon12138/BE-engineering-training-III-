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
	ID          string `json:"id" binding:"required"`
	PayID       string `json:"payId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Telephone   string `json:"telephone"`
	Description string `json:"description"`
}

type DeleteUserRequest struct {
	ID string `json:"id" binding:"required"`
}
