package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Println("making payment using razorPay", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fakePayment for testing", amount)
}

func main() {

	// razorPayGw := razorPay{}
	// paypalGw := paypal{}
	fakePaymentGw := fakePayment{}

	nwPayment := payment{
		gateway: fakePaymentGw,
	}

	nwPayment.makePayment(100)

}
