package main

import (
    //"sync"
	"log"
	"encoding/json"
	//zmqClient "github.com/yuhuichen/zmqServer/zmqClientLib"
	zmq "github.com/pebbe/zmq4"
	mongoCtr "github.com/yuhuichen/customerMgtBackendServices/controllers"
	"github.com/yuhuichen/customerMgtBackendServices/models"
)


//const zmq_broker_frontend_url_port ="tcp://localhost:5555"
const zmq_broker_backend_url_port = "tcp://localhost:5556"

func main() {

	log.Printf("Backend Service running, listening: %s\n", zmq_broker_backend_url_port)
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect(zmq_broker_backend_url_port)
	
	for {

		request, _ := responder.RecvBytes(0)
		log.Printf("Received request: [%s]\n", request)

		reply_msg, _ := query(request)
		
		log.Printf("Sent reply: [%s]\n", reply_msg)
		responder.SendBytes(reply_msg, 0)
	}

}

func query(request []byte) ([]byte, error){
	isQuery := true	
	var opt string	
	
	optKey := models.OptKey{}
	optCinfo := models.OptCinfo{}
	
	if err := json.Unmarshal(request, &optKey); err != nil {
		isQuery = false
		optCinfo := models.OptCinfo{}
		if err := json.Unmarshal(request, &optCinfo); err != nil {
			return nil, err
		}
		opt = optKey.Opt			//redundent 
	}
	
	if isQuery {
		opt = optKey.Opt
	}else{
		opt = optCinfo.Opt
	}

	reply := []byte{}
	
	cc := mongoCtr.NewCustomerController()
	
	switch {
    case opt == "INSERT":
        reply = cc.CreateCustomer(request)
	case opt == "GET":
        reply = cc.GetCustomer(request)
	case opt == "REMOVE":
        reply = cc.RemoveCustomer(request)
    }
	return reply, nil
}


