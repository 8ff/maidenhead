package main

import (
	"fmt"
	"log"

	"github.com/8ff/maidenhead"
)

func main() {
	// Convert lat/long to Maidenhead locator
	locator, err := maidenhead.GetGrid(45.5231, -122.6765)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Grid: %s\n", locator)

	// Convert Maidenhead locator to lat/long
	lat, long, err := maidenhead.GetCoordinates("CN85pm")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Coordinates: %f, %f\n", lat, long)
}
