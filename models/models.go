package models

type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	RealseDate string `json:"realseDate"`
	Runtime int `json:"runtime"`
	MPAARating string `json:"mpaa_rating"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	MovieGere []MovieGenre `json:"-"`
}

type Genre struct {
	ID int `json:"id"`
	GenreName string `json:"genre_name"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type MovieGenre struct {
	ID int `json:"id"`
	MovieID int `json:"movie_id"`
	GenreID int `json:"genre_id"`
	Genre Genre `json:"genre"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}