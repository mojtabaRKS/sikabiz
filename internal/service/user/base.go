package user

import (
	"context"
	"sikabiz/user-importer/internal/domain"

	log "github.com/sirupsen/logrus"
)

type userService struct {
	userRepository    userRepository
	addressRepository addressRepository
	logger            *log.Logger
}

type addressRepository interface {
	InsertAddress(ctx context.Context, address domain.Address) error
	GetAddressByUserId(ctx context.Context, userId string) ([]domain.Address, error)
}

type userRepository interface {
	InsertUser(ctx context.Context, user domain.User) (*domain.User, error)
	GetUser(ctx context.Context, userId string) (*domain.User, error)
}

func NewUserService(userRepo userRepository, addressRepository addressRepository, logger *log.Logger) *userService {
	return &userService{
		userRepository:    userRepo,
		addressRepository: addressRepository,
		logger:            logger,
	}
}
