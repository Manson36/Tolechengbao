package datamodels

type User struct {
	ID 			int64 `json:"id, string"`
	Username 	string `json:"username"`
}
