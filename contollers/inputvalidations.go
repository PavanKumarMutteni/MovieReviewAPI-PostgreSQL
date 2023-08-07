package contollers

import (
	entity "moviereviewapiwithdatabase/entities"
	"net/url"
)

func IsValid(movie entity.Movie) url.Values {
	errs := url.Values{}
	if movie.MovieName == "" {
		errs.Add("Movie", "The name is required!")
	}
	if movie.Director == "" {
		errs.Add("Director", "The director field is required!")
	}
	if length := len(movie.Director); !(length > 2 && length < 20) {
		errs.Add("Director", "The director field must be bewteen 2-20 chars!")
	}
	if movie.Synopsis == "" {
		errs.Add("Synopsis", "The synopsis field is required!")
	}
	if length := len(movie.Synopsis); !(length > 5 && length < 40) {
		errs.Add("Synopsis", "The synopsis field must be bewteen 5-40 chars!")
	}
	if movie.ReviewScore == 0.0 {
		errs.Add("ReviewScore", "The reviewscore is required")
	}
	if !(movie.ReviewScore >= 1 && movie.ReviewScore <= 10) {
		errs.Add("ReviewScore", "The review score must be in range 1 to 10")
	}
	if movie.YearOfRelease == 0 {
		errs.Add("YearOfRelease", "The release year field is required")
	}
	return errs
}
