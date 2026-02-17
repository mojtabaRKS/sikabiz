package entity

import (
	"fmt"
	"sikabiz/user-importer/internal/domain"
	"strings"
)

type User struct {
	Id          uint   `gorm:"primaryKey"`
	SecondaryId string `gorm:"uniqueIndex"`
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Addresses   []Address
}

func (u User) ToDomain(user User) *domain.User {
	return &domain.User{
		Id:          user.Id,
		SecondaryId: user.SecondaryId,
		Name: strings.Join([]string{
			user.FirstName,
			user.LastName,
		}, " "),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}

func (u User) FromDomain(user domain.User) (*User, error) {
	s := strings.Split(user.Name, " ")
	if len(s) < 2 {
		return nil, fmt.Errorf("name is invalid for insert")
	}

	return &User{
		SecondaryId: user.SecondaryId,
		FirstName:   s[0],
		LastName:    s[1],
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
