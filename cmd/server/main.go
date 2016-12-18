package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/oinume/go-e2e-test-sample/app"
)

func main() {
	port := 10000
	fmt.Printf("Listening on :%v\n", port)
	http.HandleFunc("/", app.Index)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalf("ListenAndServe() on :%v failed: err = %v", port, err)
	}
}
