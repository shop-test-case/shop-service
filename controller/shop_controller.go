package controller

import (
	"shop-service/model"
	"shop-service/repository"
)

type ShopController struct {
	Repo repository.IShopRepo
}

func (c *ShopController) AddShop(s *model.Shop) error {
	return c.Repo.Create(s)
}

func (c *ShopController) ListShops() ([]model.Shop, error) {
	return c.Repo.FindAll()
}
