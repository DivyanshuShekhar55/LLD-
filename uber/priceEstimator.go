package uber

import "fmt"

type PriceEstimatorService struct {
	ps       *PricingStrategy
	fareRepo *FareRepo
}

func (pes *PriceEstimatorService) estimatePrices(fareId string, products []ProductStrategy) map[ProductStrategy]float32 {
	fmt.Println("getting price estimates for all products in this src to dest")
	// this would take in the pricing strategy passed in from runtime config when the user connects and searches
	addon := pes.ps.getBasePrice()
	m := make(map[ProductStrategy]float32)
	m[&UBER_AUTO{}] = 122 + addon
	m[&UBER_XL{}] = 130 + addon
	return m
}
