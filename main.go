package main

import (
	"fmt"

	dbmock "github.com/adocoder12/golangAdvace/db_mock"
)

type App struct {
	dbmock.DBInterface
}

// Our app constructor
func NewApp(db *dbmock.DB) *App {
	return &App{
		db,
	}
}

func main() {
	db := dbmock.NewDabase("Favorites Movies Collection")
	app := NewApp(db)
	//Adding movie
	if err := app.AddMovie("life of pi", "Life philosophy", 2012); err != nil {
		fmt.Println(err.Error())
	}
	app.ShowMovies()
	//Updating movie

	fmt.Println("Updating.....")
	if err := app.UpdateMovie(1, "wolwerine", "X men life", 2015); err != nil {
		fmt.Println(err.Error())
	}
	app.ShowMovies()
	//deleting movie
	fmt.Println("deleting.....")
	app.DeleteMovie(1)
	fmt.Println("Loading movies.....")
	app.ShowMovies()
}
