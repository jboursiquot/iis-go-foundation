package main

import "fmt"

type itinerary struct {
	stops    []*airport
	minStops int
}

func newItinerary() *itinerary {
	return &itinerary{
		stops:    make([]*airport, 0),
		minStops: 2,
	}
}

func genItinerary(userCodes []string, airports map[string]*airport) (*itinerary, error) {
	it := newItinerary()

	// ensure there are enough stops in the travel plan
	if len(userCodes) < it.minStops {
		return nil, &errNotEnoughStops{stops: userCodes}
	}

	// assemble itinerary
	for _, uc := range userCodes {
		ap, exists := airports[uc]
		if !exists {
			return nil, &errUnknownCode{code: uc}
		}
		it.stops = append(it.stops, ap)
	}

	return it, nil
}

func printItinerary(it *itinerary) {
	numStops := len(it.stops)
	for i := 0; i < numStops; i++ {
		if i != numStops-1 {
			fmt.Printf("%s to ", it.stops[i].name)
		} else {
			fmt.Printf("%s\n", it.stops[i].name)
		}
	}
}
