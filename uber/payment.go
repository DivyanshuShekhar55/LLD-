package uber

import "fmt"

type PaymentStrategy interface {
	pay() bool // returns success or failure during payment
}

type UPI_CARD struct {
	card_num string
	cvv      uint16
}

func (*UPI_CARD) pay() bool {
	fmt.Println("paying with upi")
	return true
}

type CASH struct{}

func (*CASH) pay() bool {
	fmt.Println("paying in cash")
	return true
}
