package main

import (
	"emv-processor/internal/infra/emv"
	"emv-processor/internal/infra/gateway"
	httpinfra "emv-processor/internal/infra/http"
	jsonrepo "emv-processor/internal/infra/repository/json"
	"emv-processor/internal/usecase"
)

func main() {
	parser := emv.NewTLVParser()
	gateway := gateway.NewMockPaymentGateway()
	repo := jsonrepo.NewTransactionRepository("transactions.json")

	uc := usecase.NewProcessTransactionUseCase(parser, gateway, repo)

	handler := httpinfra.NewTransactionHandler(uc)
	httpinfra.StartServer(handler)
}
