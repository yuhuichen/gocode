package main

import (

	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtMicroServices/controllers"

)

const serverURI = "localhost:5000"
func main() {
	// Instantiate a new router
	r := httprouter.New()

	cc := controllers.NewCustomerController()

	r.POST("/customer", cc.CreateCustomer)

	http.ListenAndServe(serverURI, r)
}

