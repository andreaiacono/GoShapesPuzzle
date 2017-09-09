package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/cairo"
	"errors"
)

var grey = Color{0.2, 0.2, 0.2}

func openFile() (Puzzle, error) {

	filename := getFilenameFromUser()

	var puzzle Puzzle
	if filename != "" {
		puzzle, err := ReadFile(filename)
		if err != nil {
			return puzzle, err
		}
		return puzzle, nil
	}
	return puzzle, errors.New("User canceled action")
}

func getFilenameFromUser() string {
	fileChooser, err := gtk.FileChooserDialogNewWith2Buttons(
		"Open Shape Model",
		nil,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Cancel", gtk.RESPONSE_DELETE_EVENT,
		"Open", gtk.RESPONSE_ACCEPT)

	if err != nil {
		return ""
	}

	// filter for models
	filter, _ := gtk.FileFilterNew()
	filter.AddPattern("*.model")
	filter.SetName("Shape Models")
	fileChooser.AddFilter(filter)

	switcher := fileChooser.Run()
	filename := fileChooser.GetFilename()
	fileChooser.Destroy()

	// if the user pressed another button other than OK
	if switcher != -3 {
		return ""
	}

	return filename
}

func drawGrid(puzzle Puzzle, size float64, cr *cairo.Context) {

	colors := GenerateColors(len(puzzle.Pieces))

	grid := puzzle.Grid
	cellSize := size/float64(puzzle.MaxSize) - 1

	// draws all the cells
	var i, j int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				drawCell(i, j, cellSize, colors[grid[i][j]-1], cr)
			}
		}
	}
}

func drawCell(i int, j int, cellSize float64, color Color, cr *cairo.Context) {

	// computes where to locate the cell
	x := 9 + cellSize*float64(i)
	y := 9 + cellSize*float64(j)

	// draws the cell
	setColor(color, cr)
	cr.Rectangle(x, y, cellSize, cellSize)
	cr.Fill()

	// draws the border of the cell
	setColor(grey, cr)
	DrawRectangle(x, y, cellSize, cellSize, cr)
}

func setColor(color Color, context *cairo.Context) {
	context.SetSourceRGB(color.R, color.G, color.B)
}

func CreateAndStartGui(puzzle Puzzle) {

	gtk.Init(nil)

	isSolving := false
	size := float64(150)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Grid Example")
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

	statusBar, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create statusbar:", err)
	}
	statusBar.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	statusBar.SetBorderWidth(5)

	adj, err := gtk.AdjustmentNew(100.0, 0.0, 250.0, 5.0, 0.0,0.0)
	if err != nil {
		log.Fatal("Unable to create adjustement:", err)
	}
	scale, err := gtk.ScaleNew(gtk.ORIENTATION_HORIZONTAL, adj)
	if err != nil {
		log.Fatal("Unable to create scale:", err)
	}
	scale.SetHExpand(true)
	scale.Connect("value-changed", func() {
		puzzle.Speed = int(scale.GetValue())
	})

	btn, err := gtk.ButtonNewWithLabel("Find new solution")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	infoLabel, err := gtk.LabelNew("   Speed")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	btn.Connect("clicked", func() {
		if isSolving == false {
			btn.SetLabel("     Stop Finding     ")
			puzzle.Computing = true
			go solver(&puzzle, win)
		} else {
			btn.SetLabel("Find new solution")
			puzzle.Computing = false
		}
		isSolving = !isSolving
	})

	da.SetHExpand(true)
	da.SetVExpand(true)


	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		width := float64(da.GetAllocatedWidth())
		height := float64(da.GetAllocatedHeight())

		if width > height {
			size = height - 10
		} else {
			size = width - 10
		}

		// draws background
		cr.SetSourceRGB(1, 1, 1)
		cr.Rectangle(0, 0, width, height)
		cr.Fill()

		// draws border
		cr.SetSourceRGB(0.1, 0.1, 0.1)
		DrawRectangle(0, 0, width, height, cr)

		// draws the gtkGrid
		drawGrid(puzzle, size, cr)
	})

	// creates menu
	menuBar, err := gtk.MenuBarNew()
	if err != nil {
		log.Fatal("Unable to create menubar:", err)
	}

	fileMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create menu:", err)
	}

	fileMenuItem, err := gtk.MenuItemNewWithLabel("File")
	if err != nil {
		log.Fatal("Unable to create menuitem:", err)
	}

	openMenuItem, err := gtk.MenuItemNewWithLabel("Open")
	if err != nil {
		log.Fatal("Unable to create menuitem:", err)
	}

	openMenuItem.Connect("activate", func() {
		newPuzzle, err := openFile()
		if err != nil {
			log.Print(err)
			return
		}

		puzzle = newPuzzle
	})

	fileMenuItem.SetSubmenu(fileMenu)
	fileMenu.Append(openMenuItem)
	menuBar.Append(fileMenuItem)

	gtkGrid.Attach(menuBar, 0, 0, 200, 200)

	gtkGrid.Add(da)
	statusBar.Add(btn)
	statusBar.Add(infoLabel)
	statusBar.Add(scale)
	gtkGrid.Add(statusBar)


	win.Add(gtkGrid)
	win.ShowAll()

	gtk.Main()
}
