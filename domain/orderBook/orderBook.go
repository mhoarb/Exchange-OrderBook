package orderBook

import (
	"Exchange-OrderBook/domain"
)

type iOrderBook interface {
	AddOrder(order domain.Order)
	MatchOrders()
	RemoveOrder(order domain.Order)
	String() string
}

type OrderBook struct {
	BuyOrders  []domain.Order
	SellOrders []domain.Order
}
