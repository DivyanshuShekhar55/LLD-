package uber

import (
	"fmt"
)

type Rider struct {
	id int
	// name blah blah
}

type Driver struct {
	id      int
	vehicle *Vehicle
}

type Vehicle struct {
	id       int
	products []ProductStrategy
}

type Location struct {
	lat float32
	lon float32
}

type RideService struct {
	id     int
	rider  *Rider
	driver *Driver
	src    *Location
	dest   *Location
}

func main() {
	fmt.Println("hello")
}
