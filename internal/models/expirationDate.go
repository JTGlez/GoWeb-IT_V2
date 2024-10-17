package models

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrorTimeParse = errors.New("invalid expiration date format or invalid date. Use DD/MM/YYYY")
)

// ExpirationDate is a value object/wrapper to validate and store safely dates on the Product and ProductResponse structs.
type ExpirationDate struct {
	value time.Time
}

// String representation of the ExpirationDate wrapper.
func (e *ExpirationDate) String() string {
	return e.value.Format("02/01/2006")
}

func (e *ExpirationDate) Value() time.Time {
	return e.value
}

// MarshalJSON is used to create a JSON representation of ExpirationDate.
func (e *ExpirationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON is used to decode a JSON representation of a valid ExpirationDate.
func (e *ExpirationDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	expiration, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return ErrorTimeParse
	}

	e.value = expiration
	return nil
}
