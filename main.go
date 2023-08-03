package main

import (
	controller "moviereviewapiwithdatabase/contollers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getAllMovies", controller.GetAllMoviesData)
	router.GET("/getMovieByDirectorName/:director", controller.GetMovieByDirectorName)
	router.GET("/getMovieByYearOfRelease/:yearofrelease", controller.GetMovieByYearOfRelease)
	router.POST("/createMovie", controller.AddMovieRecord)
	router.DELETE("/deleteMovieByEntryNo/:entryNo", controller.DeleteMovieByEntryNo)
	router.Run("localhost:8082")
}
