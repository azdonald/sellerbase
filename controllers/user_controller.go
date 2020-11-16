package controllers

import (
	"net/http"

	"github.com/azdonald/sellerbase/src/db"

	m "github.com/azdonald/sellerbase/src/models"

	restful "github.com/emicklei/go-restful/v3"
)

//FindUser returns a User
func FindUser(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	var user m.User
	db.DB.First(&user, "id=?", id)
	res.WriteEntity(user)
}

// CreateUser is handler for creating a new user
func CreateUser(req *restful.Request, res *restful.Response) {
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

	result := db.DB.Create(&user)
	if result.Error != nil {
		res.WriteError(http.StatusBadRequest, result.Error)
	}

	//Send new registration event
	res.WriteHeaderAndEntity(http.StatusCreated, user)
}
