package main

import (
	"fmt"
	"os"

	"github.com/Gav1nnn/CineMind/Server/CineMindServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		fmt.Println("failed to set trusted proxies", err)
		return
	}

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello, cinemind")
	})

	routes.SetupProtectedRoutes(router)
	routes.SetupPublicRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Println("failed to start server", err)
	}
}
