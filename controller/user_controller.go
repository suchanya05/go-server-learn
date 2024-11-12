package controller

import (
	"encoding/json"
	"go-server/models"
	"go-server/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserService service.UserService
}

// CreateUser - สร้างผู้ใช้
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetAllUsersHandler - API endpoint สำหรับดึงข้อมูลผู้ใช้ทั้งหมด
func (uc *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUserByIDHandler - API endpoint สำหรับดึงข้อมูลผู้ใช้ตาม ID
func (uc *UserController) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)            // ใช้ gorilla/mux เพื่อดึงค่าจาก URL
	idStr := vars["id"]            // ดึง ID จาก URL
	id, err := strconv.Atoi(idStr) // แปลงเป็น int
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	users, err := uc.UserService.GetUsers(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUserByUsernameHandler - API endpoint สำหรับดึงข้อมูลผู้ใช้ตาม username
func (uc *UserController) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)          // ใช้ gorilla/mux เพื่อดึงค่าจาก URL
	username := vars["username"] // ดึง username จาก URL

	users, err := uc.UserService.GetUserName(&username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// UpdateUser - แก้ไขผู้ใช้
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uc.UserService.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser - ลบผู้ใช้
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// userIDStr := r.URL.Query().Get("id")
	// userID, err := strconv.Atoi(userIDStr)
	vars := mux.Vars(r) // ใช้ gorilla/mux เพื่อดึงค่าจาก URL
	idStr := vars["id"]
	userID, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if err := uc.UserService.DeleteUser(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
