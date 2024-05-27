package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TropicalDog17/alert-crud/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddAlert(c *gin.Context) {

	// AddAlert is a handler for adding an alert
	conditionValue, ok := model.Condition_value[c.Query("condition")]
	if !ok {
		c.JSON(400, gin.H{
			"message": "invalid condition",
		})
		return

	}

	price, err := strconv.ParseFloat(c.Query("value"), 32)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if c.Query("userId") == "" || c.Query("symbol") == "" {
		c.JSON(400, gin.H{
			"message": "missing required fields",
		})
		return
	}
	req := &model.CreateAlertRequest{
		UserId:    c.Query("userId"),
		Symbol:    c.Query("symbol"),
		Price:     float32(price),
		Condition: model.Condition(conditionValue),
	}

	err = h.GetDB().AddAlert(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "alert added",
		"alert":   req,
	})
}

func (h *Handler) GetAlert(c *gin.Context) {
	// GetAlert is a handler for getting an alert
	if c.Query("alertId") == "" {
		c.JSON(400, gin.H{
			"message": "missing required fields",
		})
		return
	}
	req := &model.GetAlertRequest{
		Id: c.Query("alertId"),
	}
	alert, err := h.GetDB().GetAlert(context.Background(), req)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "alert not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"alert": alert,
	})
}

func (h *Handler) GetAllAlerts(c *gin.Context) {
	// GetAllAlerts is a handler for getting all alerts
	alerts, err := h.GetDB().GetAllAlerts(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"alerts": alerts,
	})
}
func (h *Handler) GetAllUserAlerts(c *gin.Context) {
	// GetAllUserAlerts is a handler for getting all alerts of a user
	if c.Param("userID") == "" {
		c.JSON(400, gin.H{
			"message": "missing userId in the request",
		})
		return
	}
	fmt.Println("userID: ", c.Param("userID"))
	alerts, err := h.GetDB().GetAllUserAlerts(context.Background(), c.Param("userID"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"alerts": alerts,
	})
}

func (h *Handler) UpdateAlert(c *gin.Context) {
	// UpdateAlert is a handler for updating an alert
	if c.Query("alertId") == "" {
		c.JSON(400, gin.H{
			"message": "missing alertId in the request",
		})
		return
	}

	var req = &model.UpdateAlertRequest{}
	req.Id = c.Query("alertId")

	if c.Query("condition") != "" {
		conditionValue, ok := model.Condition_value[c.Query("condition")]
		if !ok {
			c.JSON(400, gin.H{
				"message": "invalid condition",
			})
			return

		}
		req.Condition = model.Condition(conditionValue)
	}

	if c.Query("price") != "" {
		price, err := strconv.ParseFloat(c.Query("price"), 32)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid price value",
			})
			return
		}
		req.Price = float32(price)
	}
	if req.Condition == 0 && req.Price == 0 {
		c.JSON(400, gin.H{
			"error": "at least one of condition or price must be provided",
		})
		return
	}

	err := h.GetDB().UpdateAlert(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "alert updated",
		"alert":   req,
	})
}

func (h *Handler) DeleteAlert(c *gin.Context) {
	// DeleteAlert is a handler for deleting an alert
	if c.Query("alertId") == "" {
		c.JSON(400, gin.H{
			"message": "missing alertId in the request",
		})
		return
	}

	req := &model.DeleteAlertRequest{
		Id: c.Query("alertId"),
	}
	err := h.GetDB().DeleteAlert(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "alert deleted",
	})
}

func (h *Handler) DeleteAllAlerts(c *gin.Context) {
	// DeleteAllAlerts is a handler for deleting all alerts
	err := h.GetDB().DeleteAllAlerts(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "all alerts deleted",
	})
}

func (h *Handler) HealthCheck(c *gin.Context) {
	// HealthCheck is a handler for checking the health of the service
	c.JSON(200, gin.H{
		"message": "service is healthy",
	})
}
