package uber

import (
	"fmt"
	"time"
)

type FareRepo struct {
}

func (fr *FareRepo) saveFare(fareId string, riderId string, ttl time.Time) {
	fmt.Println("saving fare to redis with a ttl")
}

func (f *FareRepo) getFare(fareId string) Fare {
	fmt.Println("getting fare from redis")
	return Fare{
		id:            fareId,
		src:           &Location{lat: 123, lon: 234},
		dest:          &Location{lat: 123, lon: 234},
		estimatedFare: 100,
		createdAt:     time.Now(),
	}
}
