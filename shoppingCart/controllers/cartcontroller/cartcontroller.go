package cartcontroller

import (
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/cart/index.html")
	tmp.Execute(response, nil)
}

func Buy(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}