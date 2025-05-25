package utils

import (
    "net/http"
)

func SetCookie(w http.ResponseWriter, name, value string) {
    http.SetCookie(w, &http.Cookie{
        Name:     name,
        Value:    value,
        Path:     "/",
        HttpOnly: true,
    })
}

func GetRoleFromCookie(r *http.Request) string {
    cookie, err := r.Cookie("role")
    if err != nil {
        return ""
    }
    return cookie.Value
}
