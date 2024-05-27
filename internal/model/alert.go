package model

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func NewAlert(userID, symbol string, value float32, condition Condition) *Alert {
	if condition == Condition_CONDITION_UNSPECIFIED {
		return nil
	}
	return &Alert{
		// ensure unique id
		Id:        GetId(userID, symbol, condition),
		UserId:    userID,
		Symbol:    symbol,
		Value:     value,
		Condition: condition,
	}
}

func UnmarshalProtoAlert(alert []byte) (*Alert, error) {
	var a = &Alert{}
	err := proto.Unmarshal(alert, a)
	if err != nil {
		return &Alert{}, err
	}
	return a, nil
}

func MarshalProtoAlert(alert *Alert) ([]byte, error) {
	data, err := proto.Marshal(alert)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetId(userID, symbol string, condition Condition) string {
	return uuid.New().String()
}

type Alerts struct {
	Alerts []*Alert `json:"alerts"`
}
