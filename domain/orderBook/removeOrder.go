package orderBook

import (
	"Exchange-OrderBook/domain/order"
)

func (ob *OrderBook) RemoveOrder(order order.Order) {
	if order.BuyOrSell == buy {
		for i, o := range ob.BuyOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.BuyOrders = append(ob.BuyOrders[:i], ob.BuyOrders[i+1:]...)

				break
			}
		}
	} else if order.BuyOrSell == Sell {
		for i, o := range ob.SellOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.SellOrders = append(ob.SellOrders[:i], ob.SellOrders[i+1:]...)
				break
			}
		}
	}

}
