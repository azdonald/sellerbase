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
	user := new(m.User)

	err := req.ReadEntity(&user)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}

	_, err = user.IsValid()
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		res.WriteError(http.StatusBadRequest, result.Error)
		return
	}

	//Send new user added event
	res.WriteHeaderAndEntity(http.StatusCreated, user)
}
