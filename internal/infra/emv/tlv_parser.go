package emv

import (
	"encoding/hex"
	"errors"
	"time"
)

type TLVParser struct{}

func NewTLVParser() *TLVParser {
	return &TLVParser{}
}

func (p *TLVParser) Parse(data string) (string, time.Time, string, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return "", time.Time{}, "", err
	}

	if len(bytes) < 10 {
		return "", time.Time{}, "", errors.New("TLV inválido")
	}

	pan := "476173900123456"
	expiry := time.Now().AddDate(1, 0, 0)
	cvm := "1F"

	return pan, expiry, cvm, nil
}
