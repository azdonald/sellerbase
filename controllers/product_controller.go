package controllers


import (
	"net/http"

	"github.com/azdonald/sellerbase/src/db"

	m "github.com/azdonald/sellerbase/src/models"

	restful "github.com/emicklei/go-restful/v3"
)


// FindProduct returns a product by the specified id
func FindProduct(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	var product m.Product
	db.DB.First(&product,"id=?" ,id)
	res.WriteHeaderAndEntity(http.StatusOK, product)
}

// GetProducts returns all products for  merchant
func GetProducts(req *restful.Request, res *restful.Response) {
	
}

// CreateProduct allows a merchant to add a new product
func CreateProduct(req *restful.Request, res *restful.Response) {
	product := new(m.Product)

	err := req.ReadEntity(&product)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}

	result := db.DB.Create(&product)
	if result.Error != nil {
		res.WriteError(http.StatusBadRequest, result.Error)
	}

	//Send new product added event
	res.WriteHeaderAndEntity(http.StatusCreated, product)
}
