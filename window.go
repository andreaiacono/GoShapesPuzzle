package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/cairo"
)

var grey = Color{0.2, 0.2, 0.2}

var drawingArea *gtk.DrawingArea

func loadFile() {

	chooser, err := gtk.FileChooserButtonNew("Choose a file", gtk.FILE_CHOOSER_ACTION_OPEN)
	if err != nil {
		log.Printf("Error opening file open dialog: %s", err)
	}

	// Show the ColorChooserDialog
	chooser.ShowNow()

	log.Printf("file: %s", chooser.GetFilename())
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
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawingarea:", err)
	}

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

	gtkGrid.Add(da)
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
		loadFile()
	})

	gtkGrid.Attach(da, 1, 1, 1, 2)
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

	menuBar, err := gtk.MenuBarNew()
	if err != nil {
		log.Fatal("Unable to create menubar:", err)
	}

	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create menu:", err)
	}
	menu.SetName("File")

	menuItem, err := gtk.MenuItemNew()
	if err != nil {
		log.Fatal("Unable to create menuitem:", err)
	}

	menuItem.Set("Open", func() {
		log.Printf("Open called.%s", menuBar)
	})
	menu.Add(menuItem)
	//
	//win.Add(menu)
	menuBar.ShowNow()

	// Add the gtkGrid to the window, and show all widgets.
	//win.Add(menuBar)

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
	//		drawingArea.QueueDraw()
	//		println(s)
	//	}
	//}()
	gtk.Main()
}
