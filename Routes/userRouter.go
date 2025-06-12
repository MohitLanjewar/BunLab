package routes

import (
	middleware "BunLab/Middleware"
	controller "burger-shop-auth/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
