package customerMgtServices

import (
	"sync"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtWebServices/controllers"
)

func MgtService_thread(wg sync.WaitGroup, serverURI string) {
	defer wg.Done()
	// Instantiate a new router
	r := httprouter.New()
	cc := controllers.NewCustomerController()
	r.POST("/customer", cc.CreateCustomer)
	r.GET("/customer/:id", cc.GetCustomer)
	r.DELETE("/customer/:id", cc.RemoveCustomer)
	
	log.Printf("Mgt Service started: [%s]\n", serverURI)

	http.ListenAndServe(serverURI, r)
	log.Fatalln("Service crashed!")
}

