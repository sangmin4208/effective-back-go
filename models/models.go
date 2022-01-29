package models

import "time"

type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	ReleaseDate time.Time `json:"releaseDate"`
	Runtime int `json:"runtime"`
	Rating int `json:"rating"`
	MPAARating string `json:"mpaa_rating"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	MovieGere []MovieGenre `json:"-"`
}

type Genre struct {
	ID int `json:"id"`
	GenreName string `json:"genre_name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type MovieGenre struct {
	ID int `json:"id"`
	MovieID int `json:"movie_id"`
	GenreID int `json:"genre_id"`
	Genre Genre `json:"genre"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}