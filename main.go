package main

import (
	"log"
	"path/filepath"
	"github.com/gotk3/gotk3/gtk"
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

	if useGui {
		CreateAndStartGui(filename)
	} else {
		puzzle, err := ReadFile(filename, gtk.Statusbar{}, false)
		puzzle.Computing = true
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Started solving...")
		Solver(&puzzle, nil)
	}
}
