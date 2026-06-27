package main

import (
	"fmt"
	"os"

	"github.com/Gav1nnn/CineMind/Server/CineMindServer/controllers"
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

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.Addmovie())
	router.POST("/register", controllers.RegisterUser())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Println("failed to start server", err)
	}
}
