package main

import (
	"fmt"
	"github.com/xtrembaker/goflix/infrastructure/web"
	"log"
	"net/http"
)

func main() {
	startWebServer()

	fmt.Println("Seems working")
}

func startWebServer() {
	log.Printf("Serving HTTP on port 8000")
	err := http.ListenAndServe(":8000", web.NewRouter().ServeHttp())
	if err != nil {
		log.Fatal(err)
	}
}
