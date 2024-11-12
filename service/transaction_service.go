package service

import (
    "go-server/repository" // ปรับให้ตรงกับโมดูลของคุณ
	"go-server/models"// ปรับให้ตรงกับโมดูลของคุณ
)

type TransactionService struct {
    TransactionRepo repository.TransactionRepository
}

func (service *TransactionService) CreateTransaction(transaction models.UserTransaction) error {
    return service.TransactionRepo.CreateTransaction(transaction)
}

func (service *TransactionService) GetTransactions() ([]models.UserTransaction, error) {
    return service.TransactionRepo.GetTransactions()
}
