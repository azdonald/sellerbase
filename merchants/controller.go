package merchants

import (
	"fmt"

	restful "github.com/emicklei/go-restful/v3"
)

// New creates a new webservice for the merchants
func New() *restful.WebService {
	fmt.Print("merchant service")
	service := new(restful.WebService)
	service.Path("/merchants").Produces(restful.MIME_JSON)
	service.Route(service.GET("").To(FindMerchant))
	return service
}

//FindMerchant returns a merchant
func FindMerchant(req *restful.Request, res *restful.Response) {
	res.WriteEntity("Hello")
}
