package usecase

import (
	"emv-processor/internal/domain/entity"
	"testing"
	"time"
)

type mockParser struct{}

type mockGateway struct{}

type mockRepo struct{}

func (m *mockParser) Parse(data string) (string, time.Time, string, error) {
	return "4532015112830366", time.Now().AddDate(1, 0, 0), "1F", nil
}

func (m *mockGateway) Authorize(tx entity.Transaction) (bool, error) {
	return true, nil
}

func (m *mockRepo) Save(tx entity.Transaction) error {
	return nil
}

func TestProcessTransaction_Success(t *testing.T) {
	uc := NewProcessTransactionUseCase(&mockParser{}, &mockGateway{}, &mockRepo{})

	tx, err := uc.Execute("dummy")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if tx.Status != "APPROVED" {
		t.Fatalf("Status esperado APPROVED, obtido %s", tx.Status)
	}
}
