# Shapes Puzzle
is a GUI application written in Go to solve shapes puzzles:

![Shapes Puzzle Application](https://raw.githubusercontent.com/andreaiacono/andreaiacono.github.io/master/img/goshapes.gif)

## Install
This project uses the ![Gotk3](https://github.com/gotk3/gotk3) provides bindings for the GTK+3. Follow the project's instructions at https://github.com/gotk3/gotk3/wiki#installation to install it.  

## Build
After having satisfied the Gotk3 dependencies, you can compile the application with
```
go build
``` 
## Run
Once built, you can run the application with:
```
./shapes
```

The application takes two parameters:
* gui=[true|false] _(defaults to true)_ if the application is to be run with a GUI or on the command line
* filename=filename _(defaults to 'models/5x6.model')_

This is the command to launch it on command line with the `4x4.model`:
```
./shapes -gui=false -filename=models/4x4.model 
```

## Models
This repo contains some models of the puzzles. If you want to solve a new one, just create a new `.model` file; the format is simply the disposition of the pieces where every piece is specified by a number (greater then 0) or a letter. So, the puzzle can have a maximum of 35 pieces (26 case unsensitive letter and 9 numbers).

## More info
The application uses a basic backtracking algorithm to find the solutions with some branch cutting for improving performances. There's a complete explanation of the internals here: https://medium.com/p/7450d6ef9e1a .

