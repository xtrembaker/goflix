package main

import (
	"fmt"
	"github.com/xtrembaker/goflix/infrastructure/persistence/sqlite"
)

func main() {
	var movieRepository = sqlite.MovieRepositoryFactory()
	fmt.Printf("movies=%v", movieRepository.List())

	fmt.Println("Seems working")
}
