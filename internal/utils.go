package internal

import (
	handler "github.com/TropicalDog17/alert-crud/internal/handler"
	"github.com/TropicalDog17/alert-crud/internal/storage"
	"github.com/gin-gonic/gin"
)

func StartServer(storage storage.StorageInterface) {
	handler := handler.NewHandler(storage)

	router := gin.Default()
	router.GET("/alert", handler.GetAlert)
	router.POST("/alert", handler.AddAlert)
	router.GET("/alerts", handler.GetAllAlerts)
	router.PUT("/alert", handler.UpdateAlert)
	router.DELETE("/alert", handler.DeleteAlert)
	router.DELETE("/alerts", handler.DeleteAllAlerts)

	router.Run("localhost:8080")
}
