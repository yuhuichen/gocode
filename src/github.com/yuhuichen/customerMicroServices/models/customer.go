package models

import "gopkg.in/mgo.v2/bson"

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
