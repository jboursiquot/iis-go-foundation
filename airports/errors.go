package main

import "fmt"

type errNotEnoughStops struct {
	stops []string
}

func (e *errNotEnoughStops) Error() string {
	return fmt.Sprint("not enough stops for an itinerary")
}

type errUnknownCode struct {
	code string
}

func (e *errUnknownCode) Error() string {
	return fmt.Sprintf("did not recognize this airport code: %s", e.code)
}
