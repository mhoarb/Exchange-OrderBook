package order

import (
	"github.com/google/uuid"
)

type IOrder interface {
	String() string
}

type Order struct {
	OrderID   uuid.UUID
	BuyOrSell string
	Price     float64
	Quantity  int32
}
