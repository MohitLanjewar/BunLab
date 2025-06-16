package controller

import (
	"BunLab/config"
	"BunLab/models"
	"BunLab/utils"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	var user models.User
	err := config.DB.Collection("users").FindOne(r.Context(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 🔑 Generate JWT
	token, err := utils.CreateToken(user.ID.Hex(), user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// 🍪 Optionally store JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // true in production (HTTPS only)
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	// ✅ Send response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token, // Optional: send token in body if not using cookie
	})
}

