package contollers

import (
	"errors"
	database "moviereviewapiwithdatabase/Database"
	entity "moviereviewapiwithdatabase/entities"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllMoviesData(context *gin.Context) {
	var movies []entity.Movie
	db := database.ConfigureDB()
	rows, err := db.Query("select * from movies")
	if err != nil {
		return
	}
	for rows.Next() {
		var entryNo int64
		var movieName string
		var reviewScore float64
		var synopsis string
		var director string
		var yearOfRelease int64
		_ = rows.Scan(&entryNo, &movieName, &reviewScore, &synopsis, &director, &yearOfRelease)
		movies = append(movies, entity.Movie{EntryNo: entryNo, MovieName: movieName, ReviewScore: reviewScore, Synopsis: synopsis, Director: director, YearOfRelease: yearOfRelease})
	}
	if len(movies) == 0 {
		context.IndentedJSON(http.StatusOK, "No Movies data available")
		return
	}
	context.IndentedJSON(http.StatusOK, movies)
}

func GetByDirectorName(director string) ([]entity.Movie, error) {
	var movies []entity.Movie
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from movies WHERE director = $1", strings.ToLower(director))
	if err != nil {
		return nil, errors.New("No movie data found with the provided director name")
	}
	count := 0
	for rows.Next() {
		var entryNo int64
		var movieName string
		var reviewScore float64
		var synopsis string
		var director string
		var yearOfRelease int64
		count++
		_ = rows.Scan(&entryNo, &movieName, &reviewScore, &synopsis, &director, &yearOfRelease)
		movies = append(movies, entity.Movie{EntryNo: entryNo, MovieName: movieName, ReviewScore: reviewScore, Synopsis: synopsis, Director: director, YearOfRelease: yearOfRelease})
	}
	if count == 0 {
		return nil, errors.New("No movie data found with the provided director name")
	}
	return movies, nil
}

func GetMovieByDirectorName(context *gin.Context) {
	directorName := context.Param("director")
	filteredMovies, err := GetByDirectorName(directorName)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Movies Found with the provided director name"})
		return
	}
	context.IndentedJSON(http.StatusOK, filteredMovies)
}

func GetByYearOfRelease(year string) ([]entity.Movie, error) {
	var movies []entity.Movie
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from movies WHERE yearofrelease = $1", year)
	if err != nil {
		return nil, errors.New("No movie data found with the provided release year")
	}
	count := 0
	for rows.Next() {
		var entryNo int64
		var movieName string
		var reviewScore float64
		var synopsis string
		var director string
		var yearOfRelease int64
		_ = rows.Scan(&entryNo, &movieName, &reviewScore, &synopsis, &director, &yearOfRelease)
		count++
		movies = append(movies, entity.Movie{EntryNo: entryNo, MovieName: movieName, ReviewScore: reviewScore, Synopsis: synopsis, Director: director, YearOfRelease: yearOfRelease})
	}
	if count == 0 {
		return nil, errors.New("No movie data found with the provided release year")
	}
	return movies, nil
}

func GetMovieByYearOfRelease(context *gin.Context) {
	yearOfRelease := context.Param("yearofrelease")
	filteredMovies, err := GetByYearOfRelease(yearOfRelease)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, filteredMovies)
}

func AddMovieRecord(context *gin.Context) {
	var newMovie entity.Movie
	db := database.ConfigureDB()
	if err := context.BindJSON(&newMovie); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Some thing went wrong while Adding a movie"})
		return
	}
	_, err := db.Exec("Insert into movies(entryno, moviename, reviewscore, synopsis, director, yearofrelease) values($1, $2, $3, $4, $5, $6)",
		newMovie.EntryNo, newMovie.MovieName, newMovie.ReviewScore, newMovie.Synopsis, strings.ToLower(newMovie.Director), newMovie.YearOfRelease)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Some thing went wrong while Adding a movie"})
	}
	context.IndentedJSON(http.StatusCreated, newMovie)
}

func DeleteByEntryNo(entryNo int64) (string, error) {
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from movies WHERE entryno = $1", entryNo)
	if err != nil || !rows.Next() {
		return "", errors.New("No movie data found with the provided entry number")
	}
	_, err1 := db.Exec("delete from movies where entryno = $1", entryNo)
	if err1 != nil {
		return "", err1
	}
	return "Deleted Sucessfully", nil
}

func DeleteMovieByEntryNo(context *gin.Context) {
	entryNo, _ := strconv.Atoi(context.Param("entryNo"))
	msg, err := DeleteByEntryNo(int64(entryNo))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, msg)
}
