package main

type Vehicle interface {
	Start(a int, b int) (int, float64)
	Turn()
	Stop()
}

type Auto struct{}

func (a Auto) Start(x int, y int) (int, float64) {
	return 0, 0
}

func (a Auto) Turn() {}

func (a Auto) Stop() {}

type Train struct{}

func (a Train) Start() {}

func (a Train) Stop() {}

func turnLeft(v Vehicle) {}

func main() {
	var a Auto
	var b Train
	turnLeft(a)
	turnLeft(b)
}
