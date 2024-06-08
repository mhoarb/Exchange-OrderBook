package main

import (
	orderStruct "Exchange-OrderBook/domain/order"
	orderBookStruct "Exchange-OrderBook/domain/orderBook"
	"Exchange-OrderBook/internal/app"
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	yes = "y"
	no  = "n"
)

func InputFromUser() error {
	reader := bufio.NewReader(os.Stdin)
	var orderBook orderBookStruct.OrderBook

	for {
		var order orderStruct.Order
		fmt.Println("welcome Text", "please enter your order in this format (B/S,Price,Quantity)")
		orderInput, _ := reader.ReadString('\n')
		orderParts := strings.Split(strings.TrimSpace(orderInput), ",")
		if len(orderParts) != 3 {
			fmt.Println("your order is not valid  please try again")
			continue
		}
		orderId, _ := uuid.NewRandom()
		order.OrderID = orderId

		order.BuyOrSell = orderParts[0]

		price, err := strconv.ParseFloat(orderParts[1], 64)
		if err != nil {
			log.Fatal(err)

		}

		order.Price = price
		quantity, err := strconv.Atoi(orderParts[2])
		if err != nil {
			log.Fatal(err)

		}
		order.Quantity = int32(quantity)

		orderBook.AddOrder(order)
		orderBook.MatchOrders()

		{
			fmt.Println("adding another Order", "Do you want to add another order?  (y/n)")
			choiceInput, _ := reader.ReadString('\n')
			choice := strings.TrimSpace(choiceInput)

			if choice == yes {
				continue
			} else if choice == no {
				fmt.Println("print result", "do you want to print all and pretty the orderBook?(y/n")
				PrintOrderChoiceInput, _ := reader.ReadString('\n')
				PrintOrderChoice := strings.TrimSpace(PrintOrderChoiceInput)
				if PrintOrderChoice == yes {
					app.JsonPrint(orderBook)
					fmt.Println("Print OrderBook", orderBook)
					app.Md5Hash(orderBook)
					break
				}
				break
			}
		}

	}
	return nil
}
