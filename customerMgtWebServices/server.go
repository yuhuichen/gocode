package main

import (
    "sync"
	mgtService "github.com/yuhuichen/customerMgtWebServices/customerMgtServices"
	//regService "github.com/yuhuichen/customerMgtWebServices/customerRegistrationServices"
)

const RegSeviceURI = "localhost:5550"
const MgtSeviceURI = "localhost:5551"


func main() {

    var wg sync.WaitGroup

	//go regService.RegService_thread(wg, RegSeviceURI)
    //wg.Add(1)
	
	//go mgtService.MgtService_thread(wg, MgtSeviceURI)
	mgtService.MgtService_thread(wg, MgtSeviceURI)
    wg.Add(1)
	
    wg.Wait()
	ch := make(chan bool)
	ch <- true			//block the application from existing
}


/*
func main(){
	r := httprouter.New()
	cc := controllers.NewCustomerController()
	r.POST("/customer", cc.CreateCustomer)
	r.GET("/customer/:id", cc.GetCustomer)
	r.DELETE("/customer/:id", cc.RemoveCustomer)
	http.ListenAndServe("serverURI", r)
	log.Fatalln("Service crashed!")
}
*/