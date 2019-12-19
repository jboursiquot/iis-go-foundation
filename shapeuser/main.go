package main

import (
	"fmt"
	"github.com/jboursiquot/iis-go-foundation/shapes"
)

func main() {
	list := []shapes.IGeoshape{
		&shapes.Rectangle{Label: "R1"},
		&shapes.Rectangle{Label: "R2"},
		&shapes.Rectangle{Label: "R3"},
		&shapes.Circle{Label: "C1"},
		&shapes.Circle{Label: "C2"},
	}
	render(list)
}

func render(list []shapes.IGeoshape) {
	for i, shape := range list {
		fmt.Printf("%d - DRAW: %v\n", i, shape.Draw())
		fmt.Printf("%d - PRNT: %v\n", i, shape.Print())
	}
}
