package main

import (
	"os"
	routes "burger-shop-auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRouter(router)

}
