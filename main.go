package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	filename, err := filepath.Abs("github.com/shapes/models/default.model")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	puzzle, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	CreateAndStartGui(puzzle)
}
