package routes

import (
    "burger-shop-auth/controllers"
    "burger-shop-auth/middleware"
    "net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/login", controllers.Login)

    http.HandleFunc("/add-user", middleware.RoleMiddleware("super_admin", controllers.AddUser))
}
