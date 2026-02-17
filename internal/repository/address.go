package repository

import (
	"context"
	"sikabiz/user-importer/internal/domain"
	"sikabiz/user-importer/internal/repository/entity"

	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *addressRepository {
	return &addressRepository{
		db: db,
	}
}

func (ar *addressRepository) InsertAddress(ctx context.Context, address domain.Address) error {
	addressModel := entity.Address{}.FromDomain(address)
	return ar.db.WithContext(ctx).Create(&addressModel).Error
}

func (ar *addressRepository) GetAddressByUserId(ctx context.Context, userId string) ([]domain.Address, error) {
	addressModels, err := gorm.G[entity.Address](ar.db).Where("user_id = ?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}

	addresses := make([]domain.Address, 0)
	for _, addressModel := range addressModels {
		addresses = append(addresses, entity.Address{}.ToDomain(addressModel))
	}

	return addresses, nil
}
