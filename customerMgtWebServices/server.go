package main

import (
    "sync"
	mgtService "github.com/yuhuichen/customerMgtWebServices/customerMgtServices"
	regService "github.com/yuhuichen/customerMgtWebServices/customerRegistrationServices"
)

const RegSeviceURI = "localhost:5550"
const MgtSeviceURI = "localhost:5551"

func main() {
    //messages := make(chan int)
    var wg sync.WaitGroup

	go regService.RegService_thread(wg, RegSeviceURI)
    wg.Add(1)
	
	go mgtService.MgtService_thread(wg, MgtSeviceURI)
    wg.Add(1)
	
    wg.Wait()
	
}