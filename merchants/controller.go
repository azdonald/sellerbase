package merchants

import (
	restful "github.com/emicklei/go-restful/v3"
)

// New creates a new webservice for the merchants
func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("merchants").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.GET("/{merchant_id}").To(FindMerchant))
	return service
}

//FindMerchant returns a merchant
func FindMerchant(req *restful.Request, res *restful.Response) {

}