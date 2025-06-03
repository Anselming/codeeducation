package main

import (
	"context"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

func main() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	if apiKey == "" || secretKey == "" {
		fmt.Println("BINANCE_API_KEY and BINANCE_SECRET_KEY must be set")
		return
	}

	client := binance.NewClient(apiKey, secretKey)

	order, err := client.NewCreateOrderService().
		Symbol("BTCUSDT").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).
		Quantity("0.001").
		Do(context.Background())

	if err != nil {
		fmt.Println("failed to create order:", err)
		return
	}

	fmt.Printf("Order ID: %d\n", order.OrderID)
}
