package drawing

import "fmt"

type drawer interface {
	Draw() string
}

// Draw takes in a drawer to draw.
func Draw(d drawer) {
	fmt.Printf("DRAW: %v\n", d.Draw())
}
