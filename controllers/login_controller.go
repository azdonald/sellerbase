package controllers

import (
	"encoding/json"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// Login handles user authentication
func Login(req *restful.Request, res *restful.Response) {
	var loginRequest map[string]interface{}
	err := json.NewDecoder(req.Request.Body).Decode(&loginRequest)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}

	//Send email with url to complete login process from here
	res.WriteEntity(loginRequest["email"])
}
