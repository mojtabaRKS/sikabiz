package domain

type User struct {
	Id          uint      `json:"id"`
	SecondaryId string    `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Addresses   []Address `json:"addresses"`
}
