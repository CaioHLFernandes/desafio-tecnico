package valueobject

import "errors"

var supportedCVMs = map[string]bool{
	"1F": true,
}

func ValidateCVM(cvm string) error {
	if !supportedCVMs[cvm] {
		return errors.New("CVM não suportado")
	}
	return nil
}
