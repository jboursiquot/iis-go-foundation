package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	airports := loadAirports()

	it, err := genItinerary(os.Args[1:], airports)
	if err != nil {
		if err, ok := err.(*errNotEnoughStops); ok {
			// do something meaningful for this kind of error
			log.Println(err)
		}
		if err, ok := err.(*errUnknownCode); ok {
			// do something meaningful for this kind of error
			log.Println(err)
		}
		fmt.Printf("Error Occured: %v\n", err)
		os.Exit(1)
	}

	printItinerary(it)
}
