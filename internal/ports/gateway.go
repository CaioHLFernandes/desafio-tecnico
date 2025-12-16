package ports

import "emv-processor/internal/domain/entity"

type PaymentGateway interface {
	Authorize(tx entity.Transaction) (bool, error)
}
