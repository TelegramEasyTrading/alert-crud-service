package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/TropicalDog17/alert-crud/internal"
	"github.com/TropicalDog17/alert-crud/internal/model"
	storage "github.com/TropicalDog17/alert-crud/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestIntegration(t *testing.T) {
	s, err := storage.NewLocalRedisClient()
	require.NoError(t, err)
	defer s.GetDB().Close()
	go internal.StartServer(s)

	time.Sleep(2 * time.Second)
	mockAlerts := GetMockAlerts()
	// delete all alerts for a fresh start
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/alerts", nil)
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Test AddAlert
	queryParams := fmt.Sprintf("userId=%s&symbol=%s&value=%f&condition=%s", mockAlerts[0].UserId, mockAlerts[0].Symbol, mockAlerts[0].Value, mockAlerts[0].Condition.String())
	req, err = http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/alert?%s", queryParams), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Add more alerts
	for _, alert := range mockAlerts[1:] {
		queryParams = fmt.Sprintf("userId=%s&symbol=%s&value=%f&condition=%s", alert.UserId, alert.Symbol, alert.Value, alert.Condition)
		req, err = http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/alert?%s", queryParams), nil)
		require.NoError(t, err)
		resp, err = client.Do(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	}

	// Test GetAllAlerts
	req, err = http.NewRequest("GET", "http://localhost:8080/alerts", nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	alerts := model.Alerts{}
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	require.NoError(t, err)
	require.Len(t, alerts.Alerts, 3)

	// Test UpdateAlert
	queryParams = fmt.Sprintf("alertId=%s&value=%f&condition=%s", alerts.Alerts[0].Id, 60000.0, model.Condition_PRICE_BELOW)
	req, err = http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/alert?%s", queryParams), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Test DeleteAlert
	req, err = http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/alert?alertId=%s", alerts.Alerts[0].Id), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Check if the alert is deleted
	req, err = http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/alert?alertId=%s", alerts.Alerts[0].Id), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	// Check if actually remained alerts are 2
	req, err = http.NewRequest("GET", "http://localhost:8080/alerts", nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	alerts = model.Alerts{}
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	require.NoError(t, err)
	require.Len(t, alerts.Alerts, 2)

	// Test DeleteAllAlerts
	req, err = http.NewRequest("DELETE", "http://localhost:8080/alerts", nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Check if all alerts are deleted
	alerts = model.Alerts{}
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	require.NoError(t, err)
	require.Len(t, alerts.Alerts, 0)

}

// Get three mock alerts for testing
func GetMockAlerts() []*model.Alert {

	return []*model.Alert{
		model.NewAlert("user1", "BTC", 50000, model.Condition_PRICE_ABOVE),
		model.NewAlert("user2", "ETH", 2000, model.Condition_PRICE_BELOW),
		model.NewAlert("user3", "DOGE", 0.5, model.Condition_PRICE_ABOVE),
	}
}
