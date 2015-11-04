package models

import "gopkg.in/mgo.v2/bson"

type (	
	OptCinfo struct {
		Opt				string `json:"opt" bson:"opt"`
		*Customer			
	}		
)

type (	
	OptKey struct {
		Opt				string `json:"opt" bson:"opt"`
		Id 				bson.ObjectId 	`json:"id" bson:"_id"`			
	}		
)
