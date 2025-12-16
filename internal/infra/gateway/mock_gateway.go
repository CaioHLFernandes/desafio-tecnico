package gateway

import (
	"emv-processor/internal/domain/entity"
	"math/rand"
	"time"
)

type MockPaymentGateway struct{}

func NewMockPaymentGateway() *MockPaymentGateway {
	rand.Seed(time.Now().UnixNano())
	return &MockPaymentGateway{}
}

func (g *MockPaymentGateway) Authorize(tx entity.Transaction) (bool, error) {
	return rand.Intn(2) == 0, nil
}
