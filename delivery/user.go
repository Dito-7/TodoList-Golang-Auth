package delivery

import (
	"TodoList-Golang-Auth/middleware"
	"TodoList-Golang-Auth/repository"
	"TodoList-Golang-Auth/usecase"
	"TodoList-Golang-Auth/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type UserHandler struct {
	userUsecase   usecase.UserUsecase
	blacklistRepo repository.BlacklistRepository
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(userUsecase usecase.UserUsecase, blacklistRepo repository.BlacklistRepository) *UserHandler {
	return &UserHandler{
		userUsecase:   userUsecase,
		blacklistRepo: blacklistRepo,
	}
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	email, ok := middleware.GetUserEmailFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome, " + email,
	})
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.RegisterUser(context.Background(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.LoginUser(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := parts[1]

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = h.blacklistRepo.AddToken(r.Context(), tokenString, claims.ExpiresAt.Time)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
