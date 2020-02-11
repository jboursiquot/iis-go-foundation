package main

import (
	"github.com/jboursiquot/iis-go-foundation/shapes"
	"github.com/jboursiquot/iis-go-foundation/shapeuser2/drawing"
)

func main() {
	r1 := &shapes.Rectangle{Label: "R1"}
	c1 := &shapes.Circle{Label: "C1"}
	drawing.Draw(r1)
	drawing.Draw(c1)
}
