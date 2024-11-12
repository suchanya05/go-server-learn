package repository

import (
	"context"
	"go-server/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	DB *pgxpool.Pool
}

func (repo *TransactionRepository) CreateTransaction(transaction models.UserTransaction) error {
	_, err := repo.DB.Exec(context.Background(),
		"INSERT INTO user_transaction (user_id, trans_type, price, last_total, trans_desc) VALUES ($1, $2, $3, $4, $5)",
		transaction.UserID, transaction.TransType, transaction.Price, transaction.LastTotal, transaction.TransDesc)
	return err
}

func (repo *TransactionRepository) GetTransactions() ([]models.UserTransaction, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT * FROM user_transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.UserTransaction
	for rows.Next() {
		var transaction models.UserTransaction
		if err := rows.Scan(&transaction.TransID, &transaction.UserID, &transaction.TransType, &transaction.Price, &transaction.LastTotal, &transaction.TransDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
