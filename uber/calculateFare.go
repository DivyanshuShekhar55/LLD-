package uber

type FarePrice struct {
	ride        *RideService
	product     *ProductStrategy
	event       *PricingStrategy
	paymentType *PaymentStrategy
}

