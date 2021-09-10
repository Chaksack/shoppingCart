package productcontroller

import (
	"html/template"
	"net/http"
	"trainingProject/shoppingCart/models"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"products" : products,
	}
	tmp, _ :=template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)
}
