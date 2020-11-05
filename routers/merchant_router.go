package routers

import (
	ctr "github.com/azdonald/sellerbase/src/controllers"
	restful "github.com/emicklei/go-restful/v3"
)

// NewMerchantService creates a new webservice for the merchants
func NewMerchantService() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/merchants").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.GET("/{id}").To(ctr.FindMerchant))
	service.Route(service.POST("").To(ctr.CreateMerchant))
	return service
}
