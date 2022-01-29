package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sangmin4208/effective-back-go/models"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context())

	id,err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
	}

	app.logger.Println("id is", id)
	movie := models.Movie {
		ID: id,
		Title: "Some movie",
		Description: "Some description",
		Year: 2021,
		ReleaseDate: time.Date(2021,01,01,01,0,0,0,time.Local),
		Runtime: 100,
		Rating:5,
		MPAARating: "PG-13",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request){
	
}