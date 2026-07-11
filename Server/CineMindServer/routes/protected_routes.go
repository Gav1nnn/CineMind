package routes

import (
	"github.com/Gav1nnn/CineMind/Server/CineMindServer/controllers"
	"github.com/Gav1nnn/CineMind/Server/CineMindServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.Addmovie())
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies())
	router.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate())
}
