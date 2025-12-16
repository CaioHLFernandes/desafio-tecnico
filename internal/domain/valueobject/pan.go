package valueobject

import "errors"

func ValidatePAN(pan string) error {
	if len(pan) < 13 || len(pan) > 19 {
		return errors.New("PAN inválido")
	}

	sum := 0
	alternate := false
	for i := len(pan) - 1; i >= 0; i-- {
		n := int(pan[i] - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alternate = !alternate
	}

	if sum%10 != 0 {
		return errors.New("PAN falhou no algoritmo de Luhn")
	}
	return nil
}
