# Shapes Puzzle
is a GUI application written in Go to solve shapes puzzles:
![Shapes Puzzle Application](2https://raw.githubusercontent.com/andreaiacono/andreaiacono.github.io/master/img/shapes.gif)

## Install
This project uses the ![Gotk3](https://github.com/gotk3/gotk3) provides bindings for the GTK+3. Follow the project's instructions at https://github.com/gotk3/gotk3/wiki#installation to install it.  

## Run
After having satisfied the Gotk3 dependencies, you can run the project with
```

``` 

## Models
This repo contains some models of the puzzles. If you want to solve a new one, just create a new `.model` file; the format is simply the disposition of the pieces where every piece is specified by a number or a letter. So, the puzzle can have a maximum of 36 pieces (26 case unsensitive letter and 10 numbers).

## More info
The application uses a basic backtracking algorithm to find the solutions with some branch cutting for improving performances.