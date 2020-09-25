package models

type AddUserReqBody struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginUserReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
