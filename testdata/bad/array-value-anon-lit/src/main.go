package main

func main() {
	cells := [2]int{1, 2}
	// An anonymous struct has no type name for the temporary.
	g := struct{ cells [2]int }{cells: cells}
	_ = g
}
