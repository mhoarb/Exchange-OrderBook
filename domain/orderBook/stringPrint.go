package orderBook

import "strings"

func (ob *OrderBook) String() string {
	var result strings.Builder
	maxOrders := len(ob.BuyOrders)
	if len(ob.SellOrders) > maxOrders {
		maxOrders = len(ob.SellOrders)
	}

	for i := 0; i < maxOrders; i++ {
		if i < len(ob.BuyOrders) {
			result.WriteString(ob.BuyOrders[i].String())
		}
		result.WriteString(" | ")
		if i < len(ob.SellOrders) {
			result.WriteString(ob.SellOrders[i].String())
		}
		result.WriteString("\n")
	}
	return result.String()

}
