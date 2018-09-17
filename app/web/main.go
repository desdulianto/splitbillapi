package main

import (
	"log"
	"net/http"

	"github.com/desdulianto/splitbillapi"
)

func main() {
	http.HandleFunc("/", splitbillapi.SplitBillHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
