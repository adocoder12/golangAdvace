package main

import (
	"fmt"

	dbmock "github.com/adocoder12/golangAdvace/db_mock"
	fileDb "github.com/adocoder12/golangAdvace/fileDB"
)

type App struct {
	dbmock.DBInterface
}

// Our app constructor
func NewApp(db *dbmock.DB, fileDb *fileDb.FileDB) *App {
	return &App{
		db,
		fileDb,
	}
}

func main() {
	db := dbmock.NewDabase("Favorites Movies Collection")
	fileDb := fileDb.NewFileDB("movies.json")
	app := NewApp(db, fileDb)
	app.DBInterface.ShowMovies()

	//Adding movie
	if err := app.DBInterface.AddMovie("life of pi", "Life philosophy", 2012); err != nil {
		fmt.Println(err.Error())
	}
	app.FileDBInterface.ShowMovies()
	//Updating movie

	fmt.Println("Updating.....")
	if err := app.DBInterface.UpdateMovie(1, "wolwerine", "X men life", 2015); err != nil {
		fmt.Println(err.Error())
	}
	app.FileDBInterface.ShowMovies()
	//deleting movie
	fmt.Println("deleting.....")
	app.DBInterface.DeleteMovie(1)
	fmt.Println("Loading movies.....")
	app.DBInterface.ShowMovies()
}
