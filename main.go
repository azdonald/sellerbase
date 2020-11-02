package main

import (
	"log"
	"merchants"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

func main() {
	restful.Add(merchants.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
