package uber

type PricingStrategy interface {
	getPrice() int
}

type NIGHT_CHARGES struct{}

func (*NIGHT_CHARGES) getPrice() int {
	return 150
}

type LONG_HOLIDAYS struct{}

func (*LONG_HOLIDAYS) getPrice() int {
	return 50
}
