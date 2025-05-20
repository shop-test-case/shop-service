package handler

import (
	"net/http"
	"shop-service/controller"
	"shop-service/model"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	Ctrl *controller.ShopController
}

func (h *ShopHandler) AddShop(c *gin.Context) {
	var shop model.Shop

	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Ctrl.AddShop(&shop); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create shop"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "shop created"})
}

func (h *ShopHandler) GetShops(c *gin.Context) {
	shops, err := h.Ctrl.ListShops()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch shops"})
		return
	}

	c.JSON(http.StatusOK, shops)
}
