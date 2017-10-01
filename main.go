package main

import (
	"log"
	"path/filepath"
	"flag"
)

func main() {

	useGui := *flag.Bool("gui", true, "shows a GUI")
	file := flag.String("filename", "", "the file containing the model")
	flag.Parse()

	filename, err := filepath.Abs(*file)
	if err != nil {
		log.Fatal(err)
	}

	// reads the puzzle
	puzzle, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	//useGui = false
	//filename, err = filepath.Abs("github.com/shapes/models/6x6.model")

	if useGui {
		CreateAndStartGui(filename, puzzle)
	} else {
		puzzle.IsRunning = true
		log.Println("Started solving...")
		Solver(&puzzle)
	}
}
