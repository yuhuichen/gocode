package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/yuhuichen/customerMgtMicroServices/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customerMicroServices"
const collectionName = "customers"
const mongoServer = "mongodb://mongodb"

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
	
	c := models.Customer{}

	if err := cc.session.DB(dbName).C(collectionName).FindId(oid).One(&c); err != nil {
		log.Println(id, err.Error())
		w.WriteHeader(404)
		return
	}

	cj, _ := json.Marshal(c)
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

func main() {
	// Instantiate a new router
	r := httprouter.New()

	cc := controllers.NewCustomerController()

	r.POST("/customer", cc.CreateCustomer)
	r.GET("/customer/:id", cc.GetCustomer)
	r.DELETE("/customer/:id", cc.RemoveCustomer)

	http.ListenAndServe("serverURI", r)
}