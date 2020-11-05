package controllers

import (
	"net/http"

	"github.com/azdonald/sellerbase/src/db"

	m "github.com/azdonald/sellerbase/src/models"

	restful "github.com/emicklei/go-restful/v3"
)

//FindMerchant returns a merchant
func FindMerchant(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	var merchant m.Merchant
	result := db.DB.First(&merchant, id)
	res.WriteEntity(merchant)
}

// CreateMerchant is handler for creating a new merchant
func CreateMerchant(req *restful.Request, res *restful.Response) {
	merchant := new(m.Merchant)
	err := req.ReadEntity(&merchant)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}
	result := db.DB.Create(&merchant)
	if result.Error != nil {
		res.WriteError(http.StatusBadRequest, result.Error)
	}
	res.WriteEntity(merchant)
}
