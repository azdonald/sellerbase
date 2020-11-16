package routers

import (
	ctr "github.com/azdonald/sellerbase/src/controllers"
	restful "github.com/emicklei/go-restful/v3"
)

// NewProductService creates a new webservice for managing products
func NewProductService() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/products").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.GET("/{id}").To(ctr.FindProduct))
	service.Route(service.POST("").To(ctr.CreateProduct))
	service.Route(service.GET("").To(ctr.GetProducts))
	return service
}
