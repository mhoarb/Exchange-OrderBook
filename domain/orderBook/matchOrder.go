package orderBook

import (
	"Exchange-OrderBook/domain"
	"fmt"
	"github.com/google/uuid"
)

func (ob *OrderBook) MatchOrders() {
	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buyOrder := ob.BuyOrders[0]
		sellOrder := ob.SellOrders[0]
		if buyOrder.Price == sellOrder.Price {

			ob.RemoveOrder(buyOrder)
			ob.RemoveOrder(sellOrder)
			if buyOrder.Quantity > sellOrder.Quantity {
				remainOrder := domain.Order{OrderID: uuid.New(),
					BuyOrSell: "B",
					Price:     buyOrder.Price,
					Quantity:  buyOrder.Quantity - sellOrder.Quantity}
				ob.AddOrder(remainOrder)

			} else if sellOrder.Quantity > buyOrder.Quantity {
				remainOrder := domain.Order{OrderID: uuid.New(),
					BuyOrSell: "S",
					Price:     sellOrder.Price,
					Quantity:  sellOrder.Quantity - buyOrder.Quantity}
				ob.AddOrder(remainOrder)

			}
			fmt.Printf("Matched order , delete this buyOrder(%v) & this sellOrder(%v)\n", buyOrder, sellOrder)

		} else {
			break
		}

	}

}
