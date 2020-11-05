package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azdonald/sellerbase/src/db"
	r "github.com/azdonald/sellerbase/src/routers"
	"github.com/emicklei/go-restful/v3"
)

func main() {
	fmt.Println("Starting server ......")
	db.InitDB()
	restful.Add(r.NewMerchantService())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
