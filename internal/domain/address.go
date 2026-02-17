package domain

type Address struct {
	Id      uint   `json:"id"`
	UserId  uint   `json:"user_id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}
