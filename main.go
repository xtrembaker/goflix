package main

import (
	"fmt"
	"github.com/xtrembaker/goflix/infrastructure/persistence/sqlite"
)

func main() {
	client := sqlite.GetInstance()
	fmt.Printf("IsConnected=%v\n", client.IsConnected())

	client.Disconnect()
	fmt.Printf("IsConnected=%v\n", client.IsConnected())

	fmt.Println("Seems working")
}
