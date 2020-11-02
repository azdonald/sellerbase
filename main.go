package main

import (
	"log"
	"github.com/azdonald/sellerbase/src/merchants"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

func main() {
	restful.Add(merchants.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
