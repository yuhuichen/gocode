
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

func zmq_req(zmq_url_port string, data []byte) ([]byte, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	requester.SendBytes(data, 0)
	//requester.Send(msg, 0)
	reply, err := requester.RecvBytes(0)
	log.Println(reply, err.Error())
	return reply, err
}

func main() {
	
	//json.NewDecoder(r.Body).Decode(&c)
	jstr := `{"firstName": "Infosys", "lastName": "Inc.", "emailAddress": "infosys@test.com", "Street": "Canary Wharf", "City": "London", "Postcode": "E14 5NP"}`

	c := Customer{}	
	json.Unmarshal([]byte(jstr), &c)	
	log.Println(1, c)
	
	cj, _ := json.Marshal(c)
	log.Println(5, cj)
	
	c.Id = bson.NewObjectId()	
	log.Println(2, c)
		
	opt := OptMsg{}
	opt.Opt = "INSERT"	
	//opt.CC = []byte(jstr)
	opt.Customer = &c
	log.Println(21, opt)
	
	optj1, _ := json.Marshal(opt)
	log.Println(3, optj1)
	
	
	opt2 := OptMsg{}
	json.Unmarshal(optj1, &opt2)
	log.Println(4, opt2)
	//json.NewDecoder(r.Body).Decode(&c)


	//c2 := Customer{}
	c2 := opt2.Customer
	//json.Unmarshal([]byte(opt2.Customer), &c2)	
	log.Println(41, c2)
	
	
	zmq_req("tcp://*:5555", optj1)
	
	//cc.session.DB(dbName).C(collectionName).Insert(c)
}