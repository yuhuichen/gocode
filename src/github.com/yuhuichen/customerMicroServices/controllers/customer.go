package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMicroServices/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customerMicroServices"
const collectionName = "customers"

type (
	CustomerController struct {
		session *mgo.Session
	}
)

func NewCustomerController(s *mgo.Session) *CustomerController {
	return &CustomerController{s}
}

func (cc CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	
	c := models.Customer{}

	if err := cc.session.DB(dbName).C(collectionName).FindId(oid).One(&c); err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}

	cj, _ := json.Marshal(c)
	log.Println(cj)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cj)
}


func (cc CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	c := models.Customer{}
	json.NewDecoder(r.Body).Decode(&c)
	
	c.Id = bson.NewObjectId()
	
	cc.session.DB(dbName).C(collectionName).Insert(c)
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

	if err := cc.session.DB(dbName).C(collectionName).RemoveId(oid); err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
