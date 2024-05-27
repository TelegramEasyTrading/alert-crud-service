package storage

import (
	"context"
	"fmt"

	"github.com/TropicalDog17/alert-crud/internal/model"
)

func (s *Storage) AddAlert(ctx context.Context, req *model.CreateAlertRequest) error {
	alert := model.NewAlert(req.UserId, req.Symbol, float32(req.Price), req.Condition)
	data, err := model.MarshalProtoAlert(alert)
	if err != nil {
		return err
	}
	// store alert in redis
	_, err = s.GetDB().HSet(ctx, "alerts", alert.Id, data).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetAlert gets an alert from the database, given an alert id.
func (s *Storage) GetAlert(ctx context.Context, req *model.GetAlertRequest) (*model.Alert, error) {

	data, err := s.GetDB().HGet(ctx, "alerts", req.Id).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("log: ", data)
	alert, err := model.UnmarshalProtoAlert([]byte(data))
	if err != nil {
		return nil, err
	}
	return alert, nil
}

func (s *Storage) GetAllUserAlerts(ctx context.Context, userId string) ([]*model.Alert, error) {
	alerts := make([]*model.Alert, 0)
	data, err := s.GetDB().HGetAll(ctx, "alerts").Result()
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		alert, err := model.UnmarshalProtoAlert([]byte(v))
		if err != nil {
			return nil, err
		}
		if alert.UserId == userId {
			alerts = append(alerts, alert)
		}
	}
	return alerts, nil
}

func (s *Storage) GetAllAlerts(ctx context.Context) ([]*model.Alert, error) {
	alerts := make([]*model.Alert, 0)
	data, err := s.GetDB().HGetAll(ctx, "alerts").Result()
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		alert, err := model.UnmarshalProtoAlert([]byte(v))
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	return alerts, nil
}

func (s *Storage) UpdateAlert(ctx context.Context, req *model.UpdateAlertRequest) error {
	alert, err := s.GetAlert(ctx, &model.GetAlertRequest{Id: req.Id})
	if err != nil {
		return err
	}
	if req.Condition != model.Condition_CONDITION_UNSPECIFIED {
		alert.Condition = req.Condition
	}
	if req.Price != 0 {
		alert.Value = float32(req.Price)
	}

	data, err := model.MarshalProtoAlert(alert)
	if err != nil {
		return err
	}

	_, err = s.GetDB().HExists(ctx, "alerts", alert.Id).Result()
	if err != nil {
		return err
	}

	_, err = s.GetDB().HSet(ctx, "alerts", alert.Id, data).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteAlert(ctx context.Context, req *model.DeleteAlertRequest) error {
	_, err := s.GetDB().HDel(ctx, "alerts", req.Id).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteAllAlerts(ctx context.Context) error {
	_, err := s.GetDB().Del(ctx, "alerts").Result()
	if err != nil {
		return err
	}
	return nil
}
