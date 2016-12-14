package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/oinume/go-e2e-test-sample/app"
)

func main() {
	mux := app.NewServeMux()
	port := 10000
	fmt.Printf("Listening on :%v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != nil {
		log.Fatalf("ListenAndServe() on :%v failed: err = %v", port, err)
	}
}
