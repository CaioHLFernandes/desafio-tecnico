package httpinfra

import (
	"encoding/json"
	"net/http"
	"time"

	"emv-processor/internal/usecase"
)

type TransactionHandler struct {
	useCase *usecase.ProcessTransactionUseCase
}

func NewTransactionHandler(uc *usecase.ProcessTransactionUseCase) *TransactionHandler {
	return &TransactionHandler{useCase: uc}
}

type transactionRequest struct {
	TLV string `json:"tlv"`
}

type transactionResponse struct {
	Status    string    `json:"status"`
	PAN       string    `json:"pan"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h *TransactionHandler) Process(w http.ResponseWriter, r *http.Request) {
	var req transactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request inválida", http.StatusBadRequest)
		return
	}

	tx, err := h.useCase.Execute(req.TLV)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	resp := transactionResponse{
		Status:    tx.Status,
		PAN:       tx.PAN,
		CreatedAt: tx.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
