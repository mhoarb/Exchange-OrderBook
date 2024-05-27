package order

import "fmt"

func (o Order) String() string {
	return fmt.Sprintf("%s,%.2f,%d", o.BuyOrSell, o.Price, o.Quantity)
}
