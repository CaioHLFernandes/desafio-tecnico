package httpinfra

import (
	"net/http"
)

func StartServer(handler *TransactionHandler) {
	http.HandleFunc("/transactions", handler.Process)
	http.ListenAndServe(":8080", nil)
}
