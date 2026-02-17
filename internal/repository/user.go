package repository

import (
	"context"
	"sikabiz/user-importer/internal/domain"
	"sikabiz/user-importer/internal/repository/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) InsertUser(ctx context.Context, user domain.User) (*domain.User, error) {
	userModel, err := entity.User{}.FromDomain(user)

	if err != nil {
		return nil, err
	}

	err = ur.db.WithContext(ctx).Where("secondary_id = ?", user.SecondaryId).FirstOrCreate(&userModel).Error
	if err != nil {
		return nil, err
	}

	res := userModel.ToDomain(*userModel)
	return res, nil
}

func (ur *userRepository) GetUser(ctx context.Context, userId string) (*domain.User, error) {
	user, err := gorm.G[entity.User](ur.db).Where("id = ?", userId).First(ctx)
	if err != nil {
		return nil, err
	}

	dUser := entity.User{}.ToDomain(user)

	return dUser, nil
}
