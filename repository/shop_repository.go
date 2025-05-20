package repository

import (
	"shop-service/model"

	"gorm.io/gorm"
)

type IShopRepo interface {
	Create(shop *model.Shop) error
	FindAll() ([]model.Shop, error)
}

type ShopRepo struct{ DB *gorm.DB }

func (r *ShopRepo) Create(shop *model.Shop) error {
	return r.DB.Create(shop).Error
}

func (r *ShopRepo) FindAll() ([]model.Shop, error) {
	var shops []model.Shop

	err := r.DB.Find(&shops).Error

	return shops, err
}
