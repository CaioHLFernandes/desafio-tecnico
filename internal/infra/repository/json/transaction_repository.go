package jsonrepo

import (
	"emv-processor/internal/domain/entity"
	"encoding/json"
	"os"
)

type TransactionRepository struct {
	file string
}

func NewTransactionRepository(file string) *TransactionRepository {
	return &TransactionRepository{file: file}
}

func (r *TransactionRepository) Save(tx entity.Transaction) error {
	f, err := os.OpenFile(r.file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(tx)
}
