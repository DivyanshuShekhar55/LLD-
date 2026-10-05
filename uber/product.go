package uber

type ProductStrategy interface {
	getBasePrice() int
}

type UBER_AUTO struct{}

func (*UBER_AUTO) getBasePrice() int {
	return 100
}

type UBER_XL struct{}

func (*UBER_XL) getBasePrice() int {
	return 200
}
