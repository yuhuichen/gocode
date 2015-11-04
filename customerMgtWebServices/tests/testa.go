
package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/yuhuichen/customerMgtWebServices/models"
	zmqClient "github.com/yuhuichen/zmqServer/zmqClientLib"
	"log"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func zmq_req(zmq_url_port string, msg string) (string, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	start := time.Now()
	requester.Send(msg, 0)
	elapsed := time.Since(start)
    log.Printf("exec time: %s", elapsed)
	reply, err := requester.Recv(0)
	log.Println(reply, err.Error())
	return reply, err
}

func zmq_req_byte(zmq_url_port string, data []byte) ([]byte, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	start := time.Now()
	requester.SendBytes(data, 0)
	elapsed := time.Since(start)
    log.Printf("exec time: %s", elapsed)
	reply, err := requester.RecvBytes(0)
	//log.Println(string(reply), err.Error())
	return reply, err
}

func main() {
	
	//json.NewDecoder(r.Body).Decode(&c)
	jstr := `{"firstName": "Infosys", "lastName": "Inc.", "emailAddress": "infosys@test.com", "Street": "Canary Wharf", "City": "London", "Postcode": "E14 5NP"}`

	c := models.Customer{}	
	json.Unmarshal([]byte(jstr), &c)	
	c.Id = bson.NewObjectId()		
	//cj, _ := json.Marshal(c)
		
	opt := models.OptCinfo{}
	opt.Opt = "INSERT"	
	opt.Customer = &c	
	optj1, _ := json.Marshal(opt)
	
	//reply, _ := zmq_req_byte("tcp://localhost:5555", optj1)
	//zmq_req("tcp://localhost:5555", string(optj1))
	
	reply, _ := zmqClient.Req_bytes("tcp://localhost:5555", optj1)
	log.Println(42, reply)

	c2 := models.Customer{}
	json.Unmarshal(reply, &c2)	
	log.Println(41, c2)
}