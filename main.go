package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azdonald/sellerbase/src/merchants"

	"github.com/emicklei/go-restful/v3"
)

func main() {
	fmt.Println("Starting server ......")

	restful.Add(merchants.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
