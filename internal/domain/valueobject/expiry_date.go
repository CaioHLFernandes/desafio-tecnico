package valueobject

import (
	"errors"
	"time"
)

func ValidateExpiry(date time.Time) error {
	if date.Before(time.Now()) {
		return errors.New("cartão expirado")
	}
	return nil
}
