package ports

import "emv-processor/internal/domain/entity"

type TransactionRepository interface {
	Save(tx entity.Transaction) error
}
