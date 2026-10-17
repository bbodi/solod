package main

type Grid struct {
	cells [2]int
}

var cells = [2]int{1, 2}

// A package-level initializer cannot use a statement expression.
var g = Grid{cells: cells}

func main() {
	_ = g
}
