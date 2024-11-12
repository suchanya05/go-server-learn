package models

// User struct สำหรับตาราง user
type User struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	UserName    string `json:"user_name"`
}

// UserTransaction struct สำหรับตาราง user_transaction
type UserTransaction struct {
	TransID   int     `json:"trans_id"`
	UserID    int     `json:"user_id"`
	TransType string  `json:"trans_type"`
	Price     float64 `json:"price"`
	LastTotal float64 `json:"last_total"`
	TransDate string  `json:"trans_date"`
	TransDesc string  `json:"trans_desc"`
}
