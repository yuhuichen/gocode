package main

import (

	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMicroServices/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	cc := controllers.NewCustomerController(getSession())

	// Get a user resource
	r.GET("/customer/:id", cc.GetCustomer)

	// Create a new user
	r.POST("/customer", cc.CreateCustomer)

	// Remove an existing user
	r.DELETE("/customer/:id", cc.RemoveCustomer)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
