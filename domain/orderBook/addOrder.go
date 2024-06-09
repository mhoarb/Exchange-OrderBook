package orderBook

import (
	"Exchange-OrderBook/domain/order"
	"fmt"
	"log/slog"
	"sort"
)

const (
	buy  = "B"
	Sell = "S"
)

func (ob *OrderBook) AddOrder(order order.Order) {
	if order.Price < 999_999 && order.Quantity < 999_999_999 {
		if order.BuyOrSell == buy {
			ob.BuyOrders = append(ob.BuyOrders, order)

			sort.Slice(ob.BuyOrders, func(i, j int) bool {
				return ob.BuyOrders[i].Price < ob.BuyOrders[j].Price
			})

		} else if order.BuyOrSell == Sell {
			ob.SellOrders = append(ob.SellOrders, order)
			sort.Slice(ob.SellOrders, func(i, j int) bool {
				return ob.SellOrders[i].Price > ob.SellOrders[j].Price
			})
		} else {
			fmt.Println(order, " its uncorrected order")

		}

	} else {
		slog.Info("message", "error", "uncorrected order")

	}
}
