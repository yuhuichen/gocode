package main

import (

	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtMicroServices/controllers"

)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	cc := controllers.NewCustomerController()

	r.POST("/customer", cc.CreateCustomer)
	r.GET("/customer/:id", cc.GetCustomer)
	r.DELETE("/customer/:id", cc.RemoveCustomer)

	http.ListenAndServe("localhost:3000", r)
}

