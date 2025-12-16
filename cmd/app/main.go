package main

import (
	"emv-processor/internal/infra/emv"
	"emv-processor/internal/infra/gateway"
	"emv-processor/internal/infra/repository/json"
	"emv-processor/internal/usecase"
	"fmt"
)

func main() {
	tlvData := "5A08476173900123456F24032512319F34031F0302"

	parser := emv.NewTLVParser()
	gateway := gateway.NewMockPaymentGateway()
	repo := jsonrepo.NewTransactionRepository("transactions.json")

	uc := usecase.NewProcessTransactionUseCase(parser, gateway, repo)

	result, err := uc.Execute(tlvData)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado da transação:", result.Status)
}
