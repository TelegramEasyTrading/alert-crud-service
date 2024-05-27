package storage_test

import (
	"context"
	"testing"

	"github.com/TropicalDog17/alert-crud/internal/model"
	"github.com/TropicalDog17/alert-crud/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()

	// load env

	s, err := storage.NewLocalRedisClient()
	require.NoError(t, err)

	// Test AddAlert
	t.Run("AddAlert", func(t *testing.T) {
		req := &model.CreateAlertRequest{
			UserId:    "user1",
			Symbol:    "BTC",
			Price:     50000,
			Condition: model.Condition_PRICE_ABOVE,
		}
		s.AddAlert(ctx, req)
	})

	// Test GetAlert
	t.Run("GetAlert", func(t *testing.T) {
		req := &model.GetAlertRequest{
			Id: "user1" + "BTC",
		}
		alert, err := s.GetAlert(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, "user1", alert.UserId)
		assert.Equal(t, "BTC", alert.Symbol)
		assert.Equal(t, float32(50000), alert.Value)
		assert.Equal(t, model.Condition_PRICE_ABOVE, alert.Condition)
	})

	// // Test GetAllAlerts
	// t.Run("GetAllAlerts", func(t *testing.T) {
	// 	alerts, err := s.GetAllAlerts(ctx)
	// 	require.NoError(t, err)
	// 	assert.Len(t, alerts, 1)
	// })

	// Test UpdateAlert
	t.Run("UpdateAlert", func(t *testing.T) {
		req := &model.UpdateAlertRequest{
			Id:        "user1" + "BTC",
			Price:     60000,
			Condition: model.Condition_PRICE_BELOW,
		}
		err := s.UpdateAlert(ctx, req)
		require.NoError(t, err)

		getReq := &model.GetAlertRequest{
			Id: "user1" + "BTC",
		}
		alert, err := s.GetAlert(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, float32(60000), alert.Value)
		assert.Equal(t, model.Condition_PRICE_BELOW, alert.Condition)
	})

	// Test DeleteAlert
	t.Run("DeleteAlert", func(t *testing.T) {
		req := &model.DeleteAlertRequest{
			Id: "alert1",
		}
		err := s.DeleteAlert(ctx, req)
		require.NoError(t, err)

		_, err = s.GetAlert(ctx, &model.GetAlertRequest{Id: "alert1"})
		assert.Error(t, err)
	})
}
