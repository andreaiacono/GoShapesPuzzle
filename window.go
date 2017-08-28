package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/cairo"
)

var grey = Color{0.2, 0.2, 0.2}

var da *gtk.DrawingArea

func loadFile() ([][]int8, error) {

	//chooser, err := gtk.FileChooserButtonNew("Choose a file", gtk.FILE_CHOOSER_ACTION_OPEN)
	//if err != nil {
		return [][]int8{}, nil
	//}
	//
	//// Show the ColorChooserDialog
	//chooser.ShowNow()
	//
	//grid, err := ReadFile("")
	//if err != nil {
	//	return grid, err
	//}
	//return grid, nil
	//// Get the results from the ColorChooserDialog
	//if chooser.Run() == int(gtk.RESPONSE_OK) {
	//	// Grab the results
	//	color := chooser.GetRGBA()
	//	log.Println("Added color:", color)
	//	// Destroy the dialog and exit the program
	//	chooser.Destroy()
	//	gtk.MainQuit()
	//	// gtk actually throws an error using MainQuit()
	//	// so I'll exit the program fully with golang
	//	// may have unintended side effects though
	//	os.Exit(0)
	//}

}

func drawGrid(puzzle Puzzle, size float64, cr *cairo.Context) {

	//colors := GenerateColors(len(puzzle.Pieces))
	colors := GenerateColors(len(puzzle.Pieces))

	grid := puzzle.Grid
	cellSize := size/float64(puzzle.MaxSize) - 1

	// draws the tiles
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

	// draws the border
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
	drawingArea, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawingarea:", err)
	}
	da = drawingArea

	statusBar, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create gtkGrid:", err)
	}
	statusBar.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	statusBar.SetBorderWidth(5)

	btn, err := gtk.ButtonNewWithLabel("Find new solution")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	infoLabel, err := gtk.LabelNew("   Ready")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	gtkGrid.Add(drawingArea)
	statusBar.Add(btn)
	statusBar.Add(infoLabel)
	gtkGrid.Add(statusBar)

	btn.Connect("clicked", func() {
		if isSolving == false {
			btn.SetLabel("Stop Finding")
		} else {
			btn.SetLabel("Find new solution")
		}
		isSolving = !isSolving
		//update()
		//loadFile()
	})

	gtkGrid.Attach(drawingArea, 1, 1, 1, 2)
	drawingArea.SetHExpand(true)
	drawingArea.SetVExpand(true)
	//gr = grids[0]

	drawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		width := float64(drawingArea.GetAllocatedWidth())
		height := float64(drawingArea.GetAllocatedHeight())

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

	menuBar, err := gtk.MenuBarNew()

	menu, _ := gtk.MenuNew()
	menu.SetName("File")

	menuItem, err := gtk.MenuItemNew()
	menuItem.Set("Open", func() {

	})
	menu.Add(menuItem)
	//
	//win.Add(menu)

	// Add the gtkGrid to the window, and show all widgets.
	gtkGrid.Add(menuBar)

	win.Add(gtkGrid)
	win.ShowAll()

	//var nSets = 1
	//go func() {
	//	for {
	//		time.Sleep(1000 * time.Millisecond)
	//		s := fmt.Sprintf("Set a label %d time(s)!", nSets)
	//		//_, err := glib.IdleAdd(drawingArea, topLabel, s)
	//		if err != nil {
	//			log.Fatal("IdleAdd() failed:", err)
	//		}
	//		nSets++
	//		//gr = grids[nSets%4]
	//		da.QueueDraw()
	//		println(s)
	//	}
	//}()
	gtk.Main()
}
