package main

import (
	"log"
	"path/filepath"
	"github.com/gotk3/gotk3/gtk"
	"time"
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

	useGui = false

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
