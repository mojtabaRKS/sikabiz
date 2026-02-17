package user

import (
	"context"
	"sikabiz/user-importer/internal/domain"
)

type UserHandler struct {
	userService userService
}

type userService interface {
	GetUser(ctx context.Context, userId string) (*domain.User, error)
}

func NewUserHandler(userService userService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
