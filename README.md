![logo](media/logo.svg)
# maidenhead
This library provides an easy and convenient way to perform conversions between Maidenhead Grid Squares and latitudes and longitudes.

# Example
Ready to use example can be found in cmd/main.go

Alternatively, you can use the following code snippet:
```go
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
```