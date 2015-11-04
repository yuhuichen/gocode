package controllers

import (
	"encoding/json"
	"log"
	"github.com/yuhuichen/customerMgtBackendServices/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customerMicroServices"
const collectionName = "customers"
const mongoServer = "mongodb://localhost"

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

func (cc CustomerController) GetCustomer(request []byte) []byte{
	
	opt := models.OptKey{}
	json.Unmarshal(request, &opt)	
	id := opt.Id 

	reply := models.Customer{}
	
	if !bson.IsObjectIdHex(id) {	
		return []byte{}
	}

	oid := bson.ObjectIdHex(id)	

	if err := cc.session.DB(dbName).C(collectionName).FindId(oid).One(&reply); err != nil {
		log.Println(id, err.Error())	
		return []byte{}
	}

	replyj, _ := json.Marshal(reply)
	return replyj
}

func (cc CustomerController) CreateCustomer(request []byte) []byte{

	opt := models.OptCinfo{}
	json.Unmarshal(request, &opt)
	
	c := opt.Customer
	c.Id = bson.NewObjectId()
	log.Printf("Customer Info:%s/n", c)
	
	if err := cc.session.DB(dbName).C(collectionName).Insert(opt.Customer); err != nil{
		return nil
	}
	cj, _ := json.Marshal(c)
	return cj
}

func (cc CustomerController) RemoveCustomer(request []byte) []byte{
	
	opt := models.OptKey{}
	json.Unmarshal(request, &opt)	
	id := opt.Id
	
	reply := models.Customer{}
	
	if !bson.IsObjectIdHex(id) {		
		return []byte{}
	}

	oid := bson.ObjectIdHex(id)	

	if err := cc.session.DB(dbName).C(collectionName).RemoveId(oid); err != nil {
		log.Println(id, err.Error())		
		return []byte{}
	}	

	replyj, _ := json.Marshal(reply)
	return replyj
}
