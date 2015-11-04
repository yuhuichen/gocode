
package main
import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	//"github.com/yuhuichen/customerMgtWebServices/models"
	"log"
	zmq "github.com/pebbe/zmq4"
)

type (
	
	Customer struct {
		Id 				bson.ObjectId 	`json:"id" bson:"_id"`
		Firstname 		string        	`json:"firstName" bson:"firstName"`
		Lastname   		string        	`json:"lastName" bson:"lastName`
		EmailAddress   	string        	`json:"emailAddress" bson:"lastName"`
		Street    		string          `json:"street" bson:"street"`
		City   			string        	`json:"city" bson:"city"`
		Postcode 		string        	`json:"postcode" bson:"postcode"`
	}
		
)

type (
	
	OptMsg struct {
		Opt				string        		`json:"opt" bson:"opt"`
		*Customer
		//CC				[]byte				`json:"cc" bson:"cc"`		
	}
		
)

func zmq_reply(zmq_url_port string){
		
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect(zmq_url_port)
	var reply_msg = "got it"
	for {
		log.Printf("Waiting for requests ... \n")
		request, _ := responder.Recv(0)
		log.Printf("Received request: [%s]\n", request)

		responder.Send(reply_msg, 0)
		log.Printf("Sent reply: [%s]\n", reply_msg)
	}
}

func zmq_reply_byte(zmq_url_port string){
		
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect(zmq_url_port)
	//var reply_msg = "got it"
	for {
		log.Printf("Waiting for requests ... \n")
		request, _ := responder.RecvBytes(0)
		log.Printf("Received request: [%s]\n", request)
		
		opt := OptMsg{}
		json.Unmarshal(request, &opt)
		log.Println(4, opt)
		log.Println(4, opt.Customer)
		
		
		//c := Customer{}	
		//json.Unmarshal(opt.Customer, &c)
		cj, _ := json.Marshal(opt.Customer)
		
		reply_msg := cj
		responder.SendBytes(reply_msg, 0)
		log.Printf("Sent reply: [%s]\n", reply_msg)
	}
}

/*
func main(){
	zmq_reply_byte("tcp://localhost:5556")
}
*/