package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/azdonald/sellerbase/src/db"
	r "github.com/azdonald/sellerbase/src/routers"
	"github.com/emicklei/go-restful/v3"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting server ......")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB(os.Getenv("db_user"), os.Getenv("db_password"))
	restful.Add(r.NewMerchantService())
	restful.Add(r.NewUserService())
	restful.Add(r.NewLoginService())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
