package main

import "fmt"

type OrderStatus string

const (
	Recieved  OrderStatus = "recieved" // iota for int
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	ChangeOrderStatus(Prepared)
}
