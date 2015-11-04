package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtWebServices/models"
	zmqClient "github.com/yuhuichen/zmqServer/zmqClientLib"
	"gopkg.in/mgo.v2/bson"
)

var (
	debug bool
	zmq_broker_frontend_url_port string
 	zmq_broker_backend_url_port string
)

type (
	CustomerController struct {
		//session interface{}
	}
)

func init(){
	debug = true
	zmq_broker_frontend_url_port ="tcp://localhost:5555"
	zmq_broker_backend_url_port = "tcp://localhost:5556"
	log.Println("ZeroMQ broker ports:")
	log.Printf("Frontend: [%s]\n", zmq_broker_frontend_url_port)
	log.Printf("Backend: [%s]\n", zmq_broker_backend_url_port)
}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

func (cc CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	
	opt := models.OptKey{}
	opt.Opt = "GET"	
	opt.Id = oid
	optj, _ := json.Marshal(opt)
	
	cj, err := zmqClient.Req_bytes(zmq_broker_frontend_url_port, optj)
	
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
	//c.Id = bson.NewObjectId()
	
	opt := models.OptCinfo{}
	opt.Opt = "INSERT"	
	opt.Customer = &c
	optj, _ := json.Marshal(opt)
	
	reply, err := zmqClient.Req_bytes(zmq_broker_frontend_url_port, optj)
	if err != nil {
		log.Println(c.Id, err.Error())
		w.WriteHeader(404)
		return
	}

	cj, _ := json.Marshal(reply)
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
	
	opt := models.OptKey{}
	opt.Opt = "GET"	
	opt.Id = oid
	optj, _ := json.Marshal(opt)
	
	if _, err := zmqClient.Req_bytes(zmq_broker_frontend_url_port, optj); err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
