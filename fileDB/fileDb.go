package fileDb

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/adocoder12/golangAdvace/model"
)

// RWMutex stands for Reader/Writer Mutual Exclusion Lock.Think of it as a security guard for your database file. It has two modes:Read Mode and Writte mode

type FileDB struct {
	FilePath string
	mu       sync.RWMutex // Protects the file from concurrent read/write corruption
}

/*  ===== The Museum Analogy ====
 * RLock() / RUnlock() (Reading): Think of a museum exhibit. Dozens of people can look at the painting (read the file) at the exact same time without a problem.

 *  Lock() / Unlock() (Writing): Think of a painter coming in to physically repaint the canvas. While they are painting (writing to the file), the museum doors must be locked so nobody bumps into the wet paint or reads half-finished work.
 */

// Constructor function to initialize our FileDB
func NewFileDB(filePath string) *FileDB {
	return &FileDB{
		FilePath: filePath,
	}
}

// Helper: Read all movies from the file
func (f *FileDB) loadMovies() ([]model.Movie, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Movie{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var movies []model.Movie
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&movies); err != nil {
		return []model.Movie{}, nil
	}
	return movies, nil
}

// Helper: Save all movies back to the file as a JSON array
func (f *FileDB) saveMovies(movies []model.Movie) error {
	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(movies)
}

func (f *FileDB) ShowMovies() error {
	f.mu.RLock()
	defer f.mu.RUnlock()

	movies, err := f.loadMovies()
	if err != nil {
		return err
	}

	fmt.Println("============ Movies (from File) =================")
	if len(movies) == 0 {
		fmt.Println("No movies to display")
		return nil
	}

	for _, m := range movies {
		fmt.Printf("ID: %d | Name: %s | Published: %d | Description: %s\n", m.ID, m.Name, m.Year, m.Description)
	}
	return nil
}

func (f *FileDB) ShowMovie(movieID int) error {
	f.mu.RLock()
	defer f.mu.RUnlock()

	movies, err := f.loadMovies()
	if err != nil {
		return err
	}

	for _, m := range movies {
		if m.ID == movieID {
			fmt.Printf("Found! Name: %s | Year: %d | Description: %s\n", m.Name, m.Year, m.Description)
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", movieID)
}

func (f *FileDB) AddMovie(name, description string, year int) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if name == "" && year == 0 && description == "" {
		return fmt.Errorf("must fill out all the fields")
	}

	movies, err := f.loadMovies()
	if err != nil {
		return err
	}

	newMovie := model.Movie{
		ID:          len(movies) + 1,
		Name:        model.Capitalize(name),
		Year:        year,
		Description: description,
	}

	movies = append(movies, newMovie)
	return f.saveMovies(movies)
}

func (f *FileDB) UpdateMovie(movieID int, name, description string, year int) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if name == "" && year == 0 && description == "" {
		return fmt.Errorf("must fill at least one field to update")
	}

	movies, err := f.loadMovies()
	if err != nil {
		return err
	}

	found := false
	for i, m := range movies {
		if m.ID == movieID {
			if name != "" {
				movies[i].Name = model.Capitalize(name)
			}
			if year != 0 {
				movies[i].Year = year
			}
			if description != "" {
				movies[i].Description = description
			}
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("movie with ID %d not found", movieID)
	}

	fmt.Println("Movie updated successfully")
	return f.saveMovies(movies)
}

func (f *FileDB) DeleteMovie(movieID int) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	movies, err := f.loadMovies()
	if err != nil {
		return err
	}

	indexToDelete := -1
	for i, m := range movies {
		if m.ID == movieID {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return fmt.Errorf("movie with ID %d not found", movieID)
	}

	// Remove from slice safely
	movies = append(movies[:indexToDelete], movies[indexToDelete+1:]...)

	fmt.Println("Movie deleted successfully")
	return f.saveMovies(movies)
}
