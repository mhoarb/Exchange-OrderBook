package cmd

import (
	"Exchange-OrderBook/domain"
	"Exchange-OrderBook/domain/orderBook"
	"Exchange-OrderBook/internal/app"
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"
)

func InputFromUser() {
	reader := bufio.NewReader(os.Stdin)
	var orderBook orderBook.OrderBook

	for {
		var order domain.Order
		fmt.Println("please enter your order in this format (B/S,Price,Quantity)")
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
			fmt.Println("Do you want to add another order?  (y/n)")
			choiceInput, _ := reader.ReadString('\n')
			choice := strings.TrimSpace(choiceInput)

			if choice == "y" {
				continue
			} else if choice == "n" {
				fmt.Println("do you want to print all and pretty the orderBook?(y/n")
				PrintOrderChoiceInput, _ := reader.ReadString('\n')
				PrintOrderChoice := strings.TrimSpace(PrintOrderChoiceInput)
				if PrintOrderChoice == "y" {
					app.JsonPrint(orderBook)
					fmt.Println(orderBook)
					app.Md5Hash(orderBook)
					break
				}
				break
			}
		}

	}
}
