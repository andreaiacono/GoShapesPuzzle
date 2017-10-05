package main

import (
	"log"
	"path/filepath"
	"flag"
)

func main() {

	useGui := flag.Bool("gui", false, "shows a GUI")
	filename := flag.String("filename", "github.com/shapes/models/5x5.model", "the file containing the model")
	flag.Parse()

	file, err := filepath.Abs(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// reads the puzzle
	puzzle, err := ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	if *useGui {
		CreateAndStartGui(file, puzzle)
	} else {
		puzzle.IsRunning = true
		log.Println("Started solving...")
		Solver(&puzzle)
	}
}
