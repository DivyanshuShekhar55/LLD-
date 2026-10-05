package uber

import "fmt"

type PriceEstimatorService struct {
	ps       *PricingStrategy
	fareRepo *FareRepo
}

func (pes *PriceEstimatorService) estimatePrices(fareId string, products []ProductStrategy) map[ProductStrategy]float32 {

	m := make(map[ProductStrategy]float32)

	fmt.Println("getting price estimates for all products in this src to dest")
	// this would take in the pricing strategy passed in from runtime config when the user
	for _, p := range products {
		basePrice := p.getBasePrice()
		rule := rule()
		total := basePrice + rule.getPrice()
		m[p] = float32(total)
	}

	return m

}

// save the fare user clicked upon to redis and apply a ttl 
// after the ttl expires user will see the updated price for the ride
func saveFare(){

}