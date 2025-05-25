package middleware

import (
    "burger-shop-auth/utils"
    "net/http"
)

func RoleMiddleware(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        role := utils.GetRoleFromCookie(r)
        if role != requiredRole && role != "super_admin" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}
