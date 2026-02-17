package entity

import "sikabiz/user-importer/internal/domain"

type Address struct {
	Id      uint `gorm:"primaryKey"`
	UserId  uint
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}

func (a Address) FromDomain(address domain.Address) Address {
	return Address{
		UserId:  address.UserId,
		Street:  address.Street,
		City:    address.City,
		State:   address.State,
		ZipCode: address.ZipCode,
		Country: address.Country,
	}
}

func (a Address) ToDomain(address Address) domain.Address {
	return domain.Address{
		Id:      address.Id,
		UserId:  address.UserId,
		Street:  address.Street,
		City:    address.City,
		State:   address.State,
		ZipCode: address.ZipCode,
		Country: address.Country,
	}
}
