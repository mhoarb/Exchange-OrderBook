package orderBook

import "Exchange-OrderBook/domain"

func (ob *OrderBook) RemoveOrder(order domain.Order) {
	if order.BuyOrSell == "B" {
		for i, o := range ob.BuyOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.BuyOrders = append(ob.BuyOrders[:i], ob.BuyOrders[i+1:]...)

				break
			}
		}
	} else if order.BuyOrSell == "S" {
		for i, o := range ob.SellOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.SellOrders = append(ob.SellOrders[:i], ob.SellOrders[i+1:]...)
				break
			}
		}
	}

}
