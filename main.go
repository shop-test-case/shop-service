package main

import (
	"shop-service/config"
	"shop-service/controller"
	"shop-service/database"
	"shop-service/handler"
	"shop-service/middleware"
	"shop-service/model"
	"shop-service/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	db.AutoMigrate(&model.Shop{})

	repo := &repository.ShopRepo{DB: db}
	ctrl := &controller.ShopController{Repo: repo}
	h := &handler.ShopHandler{Ctrl: ctrl}

	r := gin.Default()
	r.Use(cors.Default())

	auth := r.Group("/")
	auth.Use(middleware.JWT(cfg.JWTSecret))
	auth.POST("/shops", h.AddShop)
	auth.GET("/shops", h.GetShops)

	r.Run(":" + cfg.Port)
}
