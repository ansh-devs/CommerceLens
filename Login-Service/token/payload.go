package token

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        string
	UserId    string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewPayload(userId string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	currentTime := time.Now()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:       tokenId.String(),
		UserId:   userId,
		IssuedAt: currentTime,
		//ExpiredAt: currentTime.Add(time.Hour * 24 * 30),
		ExpiredAt: currentTime.Add(duration),
	}
	return payload, nil
}

func (payload *Payload) IsValid() error {
	if time.Now().After(payload.ExpiredAt) {
		fmt.Println(payload.ExpiredAt)
		return errors.New("[INFO]: token has expired for user-id : " + payload.UserId)
	}
	return nil
}
