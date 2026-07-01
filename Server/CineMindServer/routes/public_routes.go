package routes

import (
	"github.com/Gav1nnn/CineMind/Server/CineMindServer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
	router.GET("/movies", controllers.GetMovies())
}
