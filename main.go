package main

import (
	"log"
	"os"
	"path/filepath"
	"github.com/gotk3/gotk3/gtk"
	"time"
)

func main() {

	filename, err := filepath.Abs("github.com/shapes/models/5x5.model")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	useGui := true

	if useGui {
		CreateAndStartGui(filename)
	} else {
		puzzle, err := ReadFile(filename, gtk.Statusbar{}, false)
		puzzle.Computing = true
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Started solving...")
		defer elapsed()()
		Solver(&puzzle, nil)
	}
}

func elapsed() func() {
	start := time.Now()
	return func() {
		log.Printf("Execution time: %v\n", time.Since(start))
	}
}
