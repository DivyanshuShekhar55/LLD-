package uber

import "time"

type Fare struct {
	id            string
	src           *Location
	dest          *Location
	estimatedFare float32
	createdAt     time.Time
}