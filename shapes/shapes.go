package shapes

import "fmt"

type Shape interface {
	Draw() string
	Print() string
}

type Rectangle struct {
	Label string
}

func (r *Rectangle) Draw() string {
	return fmt.Sprintf("%s Rectangular Drawing", r.Label)
}

func (r *Rectangle) Print() string {
	return fmt.Sprintf("%s Rectangular Printing", r.Label)
}

// Circle is a circle.
type Circle struct {
	Label string
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("%s Circular Drawing", c.Label)
}

func (c *Circle) Print() string {
	return fmt.Sprintf("%s Circular Printing", c.Label)
}
