package entity

type Movie struct {
	EntryNo       int64   `json:"entryNo"`
	MovieName     string  `json:"movieName"`
	ReviewScore   float64 `json:"reviewScore"`
	Synopsis      string  `json:"synopsis"`
	Director      string  `json:"director"`
	YearOfRelease int64   `json:"yearOfRelease"`
}
