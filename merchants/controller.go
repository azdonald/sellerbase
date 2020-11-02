package merchants

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// New creates a new webservice for the merchants
func New() *restful.WebService {
	fmt.Print("merchant service")
	service := new(restful.WebService)
	service.Path("/merchants").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	service.Route(service.GET("/{id}").To(FindMerchant))
	service.Route(service.POST("").To(CreateMerchant))
	return service
}

//FindMerchant returns a merchant
func FindMerchant(req *restful.Request, res *restful.Response) {
	res.WriteEntity("Hello")
}

// CreateMerchant is handler for creating a new merchant
func CreateMerchant(req *restful.Request, res *restful.Response) {
	merchant := new(Merchant)
	err := req.ReadEntity(&merchant)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}
	res.WriteEntity(merchant)
}
