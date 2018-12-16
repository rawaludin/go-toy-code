package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func (coordinate *coordinate) up() {
	coordinate.y++
}

type turtle struct {
	name     string
	location coordinate
}

func main() {
	benny := &turtle{name: "Benny", location: coordinate{x: 0, y: 0}}
	benny.location.up()
	fmt.Println(benny)
}
