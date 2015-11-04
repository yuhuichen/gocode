package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtWebServices/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const zmq_broker_frontend_url_port ="tcp://*:5555"
const zmq_broker_backend_url_port = "tcp://*:5566"
const dbName = "customerMicroServices"
const collectionName = "customers"
const mongoServer = "mongodb://mongodb"


type (	
	InsertMsg struct {
		Opt				string `json:"opt" bson:"opt"`
		*models.Customer			
	}		
)

type (	
	queryMsg struct {
		Opt				string `json:"opt" bson:"opt"`
		Id 				bson.ObjectId 	`json:"id" bson:"_id"`			
	}		
)

type (
	CustomerController struct {
		session *mgo.Session
	}
)

func NewCustomerController() *CustomerController {
	return &CustomerController{getSession()}
}

func getSession() *mgo.Session {

	s, err := mgo.Dial(mongoServer)		//using docker container link
	if err != nil {
		panic(err)
	}
	return s
}

func (cc CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	
	opt := queryMsg{}
	opt.Opt = "GET"	
	opt.Id = oid
	optj, _ := json.Marshal(opt)
	
	cj, err := zmq_req(zmq_broker_frontend_url_port, optj)
	
	if err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cj)
}

//send customer info to backend service via zmq
func (cc CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	c := models.Customer{}
	json.NewDecoder(r.Body).Decode(&c)
	c.Id = bson.NewObjectId()
	
	opt := InsertMsg{}
	opt.Opt = "INSERT"	
	opt.Customer = &c
	optj, _ := json.Marshal(opt)
	
	if _, err := zmq_req(zmq_broker_frontend_url_port, optj); err != nil {
		
		log.Println(c.Id, err.Error())
		w.WriteHeader(404)
		return
	}

	cj, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	
	fmt.Fprintf(w, "%s", cj)
}

func (cc CustomerController) RemoveCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	
	opt := queryMsg{}
	opt.Opt = "GET"	
	opt.Id = oid
	optj, _ := json.Marshal(opt)
	
	if _, err := zmq_req(zmq_broker_frontend_url_port, optj); err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
