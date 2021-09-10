package main

import (
	"net/http"
	"trainingProject/shoppingCart/controllers/cartcontroller"
	"trainingProject/shoppingCart/controllers/productcontroller"
)

func main () {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)

	http.HandleFunc("/cart/index", cartcontroller.Index)
	http.HandleFunc("/cart/index", cartcontroller.Index)
	http.HandleFunc("/cart/buy", cartcontroller.Buy)



	http.ListenAndServe(":3000", nil)
}
