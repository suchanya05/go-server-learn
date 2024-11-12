package controller

import (
    "context"
    "encoding/json"
    "net/http"
    "go-server/models"
	"go-server/service"
)

type TransactionController struct {
    TransactionService service.TransactionService
}

func (controller *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
    var transaction models.UserTransaction
    if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // คำนวณ last_total
    var lastTotal float64
    err := controller.TransactionService.TransactionRepo.DB.QueryRow(context.Background(),
        "SELECT COALESCE(SUM(price), 0) FROM user_transaction WHERE user_id = $1", transaction.UserID).Scan(&lastTotal)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    transaction.LastTotal = lastTotal + transaction.Price

    if err := controller.TransactionService.CreateTransaction(transaction); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (controller *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
    transactions, err := controller.TransactionService.GetTransactions()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transactions)
}
