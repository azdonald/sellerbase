package controllers

import (
	"net/http"

	"github.com/azdonald/sellerbase/src/db"

	m "github.com/azdonald/sellerbase/src/models"

	restful "github.com/emicklei/go-restful/v3"
)

type register struct {
	m.User
	m.Merchant
}

//FindMerchant returns a merchant
func FindMerchant(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	var merchant m.Merchant
	db.DB.First(&merchant,"id=?" ,id)
	res.WriteEntity(merchant)
}

// CreateMerchant is handler for creating a new merchant
func CreateMerchant(req *restful.Request, res *restful.Response) {
	merchantRegistration := new(register)

	err := req.ReadEntity(&merchantRegistration)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}

	merchant := new(m.Merchant)
	user := new(m.User)
	merchant.BusinessName = merchantRegistration.BusinessName
	user.FirstName = merchantRegistration.FirstName
	user.LastName = merchantRegistration.LastName
	user.Email = merchantRegistration.Email
	user.Merchant = *merchant

	result := db.DB.Create(&user)
	if result.Error != nil {
		res.WriteError(http.StatusBadRequest, result.Error)
	}

	//Send new registration event
	res.WriteHeaderAndEntity(http.StatusCreated, user)
}
