package dbmock

import (
	"fmt"
	"time"

	"github.com/adocoder12/golangAdvace/model"
)

type DB struct {
	Database_Name string
	Movies        []model.Movie
	CreatedAt     time.Time
}

// Our DB constructor
// Returns a pointer (*DB) so methods can actually modify this specific database
func NewDabase(database_Name string) *DB {
	return &DB{
		Database_Name: database_Name,
		Movies:        []model.Movie{},
		CreatedAt:     time.Now(),
	}
}

func (db *DB) ShowMovies() error {
	fmt.Println("============ Movies =================")
	if len(db.Movies) == 0 {
		fmt.Println()
		return fmt.Errorf("No movies to display")
	}
	for _, m := range db.Movies {
		fmt.Printf("Movie ID: %d\nName: %s\nPublished: %d\nDescription: %s\n", m.ID, m.Name, m.Year, m.Description)
		if len(db.Movies) > 1 {
			fmt.Println("***=============================***")
		}
	}
	return nil
}
func (db *DB) AddMovie(name, description string, year int) error {
	if name == "" && year == 0 && description == "" {
		return fmt.Errorf("Must fill all the files ")
	}
	newMovie := model.Movie{
		ID:          len(db.Movies) + 1,
		Name:        name,
		Year:        year,
		Description: description,
	}

	db.Movies = append(db.Movies, newMovie)
	return nil
}
func (db *DB) ShowMovie(movieID int) error {
	for _, m := range db.Movies {
		if m.ID == movieID {
			fmt.Printf("Found! Name: %s | Year: %d | Description: %s\n", m.Name, m.Year, m.Description)
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", movieID)
}

func (db *DB) UpdateMovie(movieID int, name, description string, year int) error {
	if name == "" && year == 0 && description == "" {
		return fmt.Errorf("must fill at least one field to update")
	}

	for i, m := range db.Movies {
		if m.ID == movieID {
			if name != "" {
				db.Movies[i].Name = model.Capitalize(name)
			}
			if year != 0 {
				db.Movies[i].Year = year
			}
			if description != "" {
				db.Movies[i].Description = description
			}
			fmt.Println("Movie updated successfully")
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", movieID)
}

func (db *DB) DeleteMovie(movieID int) error {
	for i, m := range db.Movies {
		if m.ID == movieID {
			// db.Movies[i:] targets everything from index 1 to the end: [2, 3, 4]
			//  db.Movies[i+1:] targets everything after index 1: [3, 4]
			// The copy function pastes [3, 4] over [2, 3, 4].
			copy(db.Movies[i:], db.Movies[i+1:])

			// Slice changes from [1, 3, 4, 4] to [1, 3, 4, (empty)]
			db.Movies[len(db.Movies)-1] = model.Movie{} // Clear last element

			// Slice changes from [1, 3, 4, (empty)] down to just [1, 3, 4].
			db.Movies = db.Movies[:len(db.Movies)-1] // Truncate slice
			fmt.Println("Movie deleted successfully")
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", movieID)
}

type DBInterface interface {
	ShowMovies() error
	ShowMovie(movieID int) error
	AddMovie(name string, description string, year int) error
	UpdateMovie(movieID int, name, description string, year int) error
	DeleteMovie(movieID int) error
}
