package main

import (
	"fmt"

	dbmock "github.com/adocoder12/golangAdvace/db_mock"
	fileDb "github.com/adocoder12/golangAdvace/fileDB"
)

type App struct {
	// Embedding the interface allows App to accept ANY database type
	// that implements dbmock.DBInterface (both dbmock and fileDB do!)
	dbmock.DBInterface
}

// Our app constructor accepts the interface, not a concrete struct pointer
func NewApp(database dbmock.DBInterface) *App {
	return &App{
		DBInterface: database,
	}
}

func main() {
	// --- POLYMORPHISM IN ACTION ---
	// Option 1: Use the in-memory RAM database
	// database := dbmock.NewDabase("Favorites Movies Collection")
	// db := dbmock.NewDabase("Favorites Movies Collection")
	fileDb := fileDb.NewFileDB("movies.json")
	// Inject whichever database you want into the App
	app := NewApp(fileDb)
	app.ShowMovies()

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
