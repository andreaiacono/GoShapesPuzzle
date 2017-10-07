package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/cairo"
	"fmt"
)

var solutionsIndex int

func ShowSolutions(puzzle Puzzle) {

	gtk.Init(nil)

	cellSize := float64(150)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Solutions")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//win.SetPosition(gtk.WIN_POS_CENTER)
	width, height := 340, 400
	win.SetDefaultSize(width, height)

	// Create a new gtkGrid widget to arrange child widgets
	gtkGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create gtkGrid:", err)
	}
	gtkGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	gtkGrid.SetBorderWidth(5)

	// Create some widgets to put in the gtkGrid.
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawingarea:", err)
	}

	controlsGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create statusbar:", err)
	}
	controlsGrid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	controlsGrid.SetBorderWidth(5)

	label, err := gtk.LabelNew(getMessage(puzzle))
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	label.SetMarginStart(10)
	solutionsIndex = 0
	btnLeft, err := gtk.ButtonNewWithLabel("<")
	if err != nil {
		log.Fatal("Unable to create button left:", err)
	}
	btnLeft.SetSensitive(false)

	btnRight, err := gtk.ButtonNewWithLabel(">")
	if err != nil {
		log.Fatal("Unable to create button right:", err)
	}
	btnRight.SetSensitive(solutionsIndex < len(*puzzle.Solutions)-1)

	btnLeft.Connect("clicked", func() {
		if solutionsIndex > 0 {
			solutionsIndex --
		}
		btnLeft.SetSensitive(solutionsIndex > 0)
		btnRight.SetSensitive(solutionsIndex < len(*puzzle.Solutions)-1)
		label.SetText(getMessage(puzzle))
		win.QueueDraw()
	})

	btnRight.Connect("clicked", func() {
		if solutionsIndex < len(*puzzle.Solutions)-1 {
			solutionsIndex ++
		}
		btnRight.SetSensitive(solutionsIndex < len(*puzzle.Solutions)-1)
		btnLeft.SetSensitive(solutionsIndex > 0)
		label.SetText(getMessage(puzzle))
		win.QueueDraw()
	})

	da.SetHExpand(true)
	da.SetVExpand(true)

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		windowWidth := float64(da.GetAllocatedWidth())
		windowHeight := float64(da.GetAllocatedHeight())
		windowRatio := windowWidth / windowHeight

		puzzleWidth := float64(len(puzzle.WorkingGrid[0]))
		puzzleHeight := float64(len(puzzle.WorkingGrid))
		puzzleRatio := puzzleWidth / puzzleHeight

		if windowRatio > puzzleRatio {
			cellSize = (windowHeight - 20) / puzzleHeight
		} else {
			cellSize = (windowWidth - 20) / puzzleWidth
		}

		// draws background
		cr.SetSourceRGB(1, 1, 1)
		cr.Rectangle(0, 0, windowWidth, windowHeight)
		cr.Fill()

		// draws border
		cr.SetSourceRGB(0.1, 0.1, 0.1)
		DrawRectangle(0, 0, windowWidth, windowHeight, cr, "")

		// draws the gtkGrid
		if len(*puzzle.Solutions) > 0 {
			drawGrid(puzzle, (*puzzle.Solutions)[solutionsIndex], cellSize, cr)
		} else {
			label.SetText("No solutions found yet")
		}
	})

	gtkGrid.Add(da)
	controlsGrid.Add(btnLeft)
	controlsGrid.Add(btnRight)
	controlsGrid.Add(label)
	gtkGrid.Add(controlsGrid)
	win.Add(gtkGrid)
	win.ShowAll()

	gtk.Main()
}

func getMessage(puzzle Puzzle) string {
	return fmt.Sprintf("Solution #%d / %d", solutionsIndex+1, len(*puzzle.Solutions))
}
