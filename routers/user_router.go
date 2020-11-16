package routers

import (
	ctr "github.com/azdonald/sellerbase/src/controllers"
	restful "github.com/emicklei/go-restful/v3"
)

//NewUserService handles user related fuctionality
func NewUserService()  *restful.WebService {
	service := new(restful.WebService)
	service.Path("/users").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.GET("/{id}").To(ctr.FindUser))
	service.Route(service.POST("").To(ctr.CreateUser))
	return service
}
