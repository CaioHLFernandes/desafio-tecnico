package usecase

import (
	"emv-processor/internal/domain/entity"
	"emv-processor/internal/domain/valueobject"
	"emv-processor/internal/ports"
	"time"
)

type TLVParser interface {
	Parse(data string) (string, time.Time, string, error)
}

type ProcessTransactionUseCase struct {
	parser  TLVParser
	gateway ports.PaymentGateway
	repo    ports.TransactionRepository
}

func NewProcessTransactionUseCase(p TLVParser, g ports.PaymentGateway, r ports.TransactionRepository) *ProcessTransactionUseCase {
	return &ProcessTransactionUseCase{p, g, r}
}

func (uc *ProcessTransactionUseCase) Execute(tlv string) (*entity.Transaction, error) {
	pan, expiry, cvm, err := uc.parser.Parse(tlv)
	if err != nil {
		return nil, err
	}

	if err := valueobject.ValidatePAN(pan); err != nil {
		return nil, err
	}
	if err := valueobject.ValidateExpiry(expiry); err != nil {
		return nil, err
	}
	if err := valueobject.ValidateCVM(cvm); err != nil {
		return nil, err
	}

	tx := entity.Transaction{
		PAN:        pan,
		ExpiryDate: expiry,
		CVM:        cvm,
		CreatedAt:  time.Now(),
	}

	approved, _ := uc.gateway.Authorize(tx)
	if approved {
		tx.Status = "APPROVED"
	} else {
		tx.Status = "DECLINED"
	}

	_ = uc.repo.Save(tx)

	return &tx, nil
}
