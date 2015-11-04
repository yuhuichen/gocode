package models

type (	
	OptCinfo struct {
		Opt				string `json:"opt" bson:"opt"`
		*Customer			
	}		
)

type (	
	OptKey struct {
		Opt				string `json:"opt" bson:"opt"`
		Id 				string 	`json:"id" bson:"_id"`			
	}		
)
