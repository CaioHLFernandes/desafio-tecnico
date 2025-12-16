package httpinfra

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"emv-processor/internal/domain/entity"
	"emv-processor/internal/usecase"
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

func TestTransactionHandler_Process(t *testing.T) {
	uc := usecase.NewProcessTransactionUseCase(
		&mockParser{},
		&mockGateway{},
		&mockRepo{},
	)

	handler := NewTransactionHandler(uc)

	body := []byte(`{"tlv":"dummy"}`)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler.Process(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Status esperado 200, obtido %d", rec.Code)
	}
}
