package main

import (
	"log"
	"net/http"

	"github.com/desdulianto/splitbillapi/handler"
)

func main() {
	http.HandleFunc("/", splitbillapi.SplitBillHandler)
	http.HandleFunc("/healthz", splitbillapi.HealthHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
