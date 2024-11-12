package repository

import (
	"context"
	"go-server/models"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

// CreateUser - สร้างผู้ใช้
func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	var newUser models.User
	err := r.DB.QueryRow(context.Background(), `
		INSERT INTO "users" (name, last_name, phone_number, user_name)
		VALUES ($1, $2, $3, $4) RETURNING user_id, name, last_name, phone_number, user_name`,
		user.Name, user.LastName, user.PhoneNumber, user.UserName).Scan(&newUser.UserID, &newUser.Name, &newUser.LastName, &newUser.PhoneNumber, &newUser.UserName)
	return newUser, err
}

// GetUsers - อ่านผู้ใช้ทั้งหมด
func (r *UserRepository) GetUsers(id *int, username *string) ([]models.User, error) {
	var users []models.User
	var args []interface{}

	// สร้าง Builder สำหรับ query
	var query strings.Builder
	query.WriteString(`SELECT user_id, name, last_name, phone_number, user_name FROM "users" WHERE 1=1`)

	// ตรวจสอบว่ามีการระบุ ID หรือไม่
	if id != nil {
		query.WriteString(` AND user_id = $1`)
		args = append(args, *id) // เพิ่ม id ใน args
	}

	// ตรวจสอบว่ามีการระบุ username หรือไม่
	if username != nil {
		if id != nil {
			query.WriteString(` AND user_name = $2`)
		} else {
			query.WriteString(` AND user_name = $1`)
		}
		args = append(args, *username)
	}

	// ดำเนินการ query
	rows, err := r.DB.Query(context.Background(), query.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.LastName, &user.PhoneNumber, &user.UserName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser - แก้ไขผู้ใช้
func (r *UserRepository) UpdateUser(user models.User) error {
	_, err := r.DB.Exec(context.Background(), `
		UPDATE "users"
		SET name = $1, last_name = $2, phone_number = $3, user_name = $4
		WHERE user_id = $5`,
		user.Name, user.LastName, user.PhoneNumber, user.UserName, user.UserID)
	return err
}

// DeleteUser - ลบผู้ใช้
func (r *UserRepository) DeleteUser(userID int) error {
	_, err := r.DB.Exec(context.Background(), `DELETE FROM "users" WHERE user_id = $1`, userID)
	return err
}
