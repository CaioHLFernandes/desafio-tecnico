package valueobject

import "testing"

func TestValidatePAN_Valid(t *testing.T) {
	pan := "4532015112830366"
	if err := ValidatePAN(pan); err != nil {
		t.Fatalf("PAN válido foi rejeitado: %v", err)
	}
}

func TestValidatePAN_Invalid(t *testing.T) {
	pan := "123456789"
	if err := ValidatePAN(pan); err == nil {
		t.Fatalf("PAN inválido foi aceito")
	}
}
