package customerRegistrationServices

import (
	"sync"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtWebServices/controllers"

)

func RegService_thread(wg sync.WaitGroup, serverURI string) {
	defer wg.Done()
	// Instantiate a new router
	r := httprouter.New()
	cc := controllers.NewCustomerController()
	r.POST("/customer", cc.CreateCustomer)
	http.ListenAndServe(serverURI, r)
	log.Fatalln("Service crashed!")
}

