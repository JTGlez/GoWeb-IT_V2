package models

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrorTimeParse = errors.New("invalid expiration date format or invalid date. Use DD/MM/YYYY")
)

type ExpirationDate struct {
	value time.Time
}

func NewExpirationDate(dateStr string) (*ExpirationDate, error) {
	expiration, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return nil, ErrorTimeParse
	}

	return &ExpirationDate{
		value: expiration,
	}, nil
}

func (e ExpirationDate) String() string {
	return e.value.Format("02/01/2006")
}

func (e ExpirationDate) Value() time.Time {
	return e.value
}

func (e ExpirationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *ExpirationDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	expiration, err := NewExpirationDate(dateStr)
	if err != nil {
		return err
	}
	e.value = expiration.value
	return nil
}
