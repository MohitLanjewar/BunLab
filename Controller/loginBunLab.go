package controllers

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

	utils.SetCookie(w, "role", user.Role)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "role": user.Role})
}
