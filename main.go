package main

import (
	"github.com/TropicalDog17/alert-crud/internal/handler"
	"github.com/TropicalDog17/alert-crud/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	s, err := storage.NewRedisClient()
	if err != nil {
		panic(err)
	}
	defer s.GetDB().Close()
	handler := handler.NewHandler(s)

	router := gin.Default()
	router.GET("/alert", handler.GetAlert)
	router.POST("/alert", handler.AddAlert)
	router.GET("/alerts", handler.GetAllAlerts)
	router.GET("/alerts/user/:userID", handler.GetAllUserAlerts)
	router.PUT("/alert", handler.UpdateAlert)
	router.DELETE("/alert", handler.DeleteAlert)
	router.DELETE("/alerts", handler.DeleteAllAlerts)

	router.Run(":8080")
}
