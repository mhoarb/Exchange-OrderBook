package orderBook

import (
	"Exchange-OrderBook/domain/order"
)

type iOrderBook interface {
	AddOrder(order order.Order)
	MatchOrders()
	RemoveOrder(order order.Order)
	String() string
}

type OrderBook struct {
	BuyOrders  []order.Order
	SellOrders []order.Order
}
