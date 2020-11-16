package routers

import (
	ctr "github.com/azdonald/sellerbase/src/controllers"
	restful "github.com/emicklei/go-restful/v3"
)

//NewLoginService handles login
func NewLoginService()  *restful.WebService {
	service := new(restful.WebService)
	service.Path("/login").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.POST("").To(ctr.Login))
	return service
}
