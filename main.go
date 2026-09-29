package main

type App struct {
}

// Our app constructor accepts the interface, not a concrete struct pointer
func NewApp() *App {
	return &App{}
}

func main() {

	// app := NewApp()

}
