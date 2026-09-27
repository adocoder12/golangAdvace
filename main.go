package main

import (
	"fmt"

	dbmock "github.com/adocoder12/golangAdvace/db_mock"
	fileDb "github.com/adocoder12/golangAdvace/fileDB"
)

type App struct {
	dbmock.DBInterface
	fileDb.FileDBInterface
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
	app.FileDBInterface.ShowMovies()

	//Adding movie
	if err := app.FileDBInterface.AddMovie("life of pi", "Life philosophy", 2012); err != nil {
		fmt.Println(err.Error())
	}
	app.FileDBInterface.ShowMovies()
	//Updating movie

	fmt.Println("Updating.....")
	if err := app.FileDBInterface.UpdateMovie(1, "wolwerine", "X men life", 2015); err != nil {
		fmt.Println(err.Error())
	}
	app.FileDBInterface.ShowMovies()
	//deleting movie
	fmt.Println("deleting.....")
	app.FileDBInterface.DeleteMovie(1)
	fmt.Println("Loading movies.....")
	app.FileDBInterface.ShowMovies()
}
