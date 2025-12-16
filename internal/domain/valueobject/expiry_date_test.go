package valueobject

import (
	"testing"
	"time"
)

func TestValidateExpiry_Valid(t *testing.T) {
	future := time.Now().AddDate(1, 0, 0)
	if err := ValidateExpiry(future); err != nil {
		t.Fatalf("Data futura foi rejeitada: %v", err)
	}
}

func TestValidateExpiry_Invalid(t *testing.T) {
	past := time.Now().AddDate(-1, 0, 0)
	if err := ValidateExpiry(past); err == nil {
		t.Fatalf("Data expirada foi aceita")
	}
}
