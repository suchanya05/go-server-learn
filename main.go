package main

import (
	"context"
	"go-server/controller" // ปรับให้ตรงกับโมดูลของคุณ
	"go-server/repository" // ปรับให้ตรงกับโมดูลของคุณ
	"go-server/service"    // ปรับให้ตรงกับโมดูลของคุณ
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbpool *pgxpool.Pool

func main() {
	// โหลดไฟล์ .env
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// เชื่อมต่อกับฐานข้อมูล
	connStr := os.Getenv("DB_URL")
	dbpool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	// สร้าง Repository
	userRepo := repository.UserRepository{DB: dbpool}
	transactionRepo := repository.TransactionRepository{DB: dbpool}

	// สร้าง Service
	userService := service.UserService{UserRepo: userRepo}
	transactionService := service.TransactionService{TransactionRepo: transactionRepo}

	// สร้าง Controller
	userController := controller.UserController{UserService: userService}
	transactionController := controller.TransactionController{TransactionService: transactionService}

	// กำหนด routes
	router := mux.NewRouter()
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")                                      // สร้างผู้ใช้
	router.HandleFunc("/api/users", userController.GetAllUsersHandler).Methods("GET")                               // ดึงข้อมูลผู้ใช้ทั้งหมด
	router.HandleFunc("/api/users/get/{id}", userController.GetUserByIDHandler).Methods("GET")                      // ดึงข้อมูลผู้ใช้ตาม ID
	router.HandleFunc("/api/users/get-username/{username}", userController.GetUserByUsernameHandler).Methods("GET") // ดึงข้อมูลผู้ใช้ตาม username
	router.HandleFunc("/api/users/update", userController.UpdateUser).Methods("POST")                               // แก้ไขผู้ใช้
	router.HandleFunc("/api/users/delete/{id}", userController.DeleteUser).Methods("GET")                           // ลบผู้ใช้

	router.HandleFunc("/api/transactions", transactionController.CreateTransaction).Methods("POST")  // สร้างธุรกรรม
	router.HandleFunc("/api/transactions/get", transactionController.GetTransactions).Methods("GET") // อ่านธุรกรรม

	// เริ่ม server
	http.ListenAndServe(":8080", router)
	log.Fatal("ListenAndServe :8080")
}
