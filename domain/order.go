package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type Order struct {
	OrderID   uuid.UUID
	BuyOrSell string
	Price     float64
	Quantity  int32
}

func (o Order) String() string {
	return fmt.Sprintf("%s,%.2f,%d", o.BuyOrSell, o.Price, o.Quantity)
}
