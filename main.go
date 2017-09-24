package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	filename, err := filepath.Abs("github.com/shapes/models/6x6.model")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	CreateAndStartGui(filename)
}
