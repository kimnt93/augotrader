package types

import "time"

type Signal struct {
	Type				string 		`json:type`
	Time           		time.Time 	`json:"time"`
	Name           		string    	`json:"name"`
	Symbol         		string    	`json:"symbol"`
	Price          		float64   	`json:"price"`
	Position           	float64    	`json:"position"`
}


type AccountConfig struct {
	Time           			time.Time 		`json:"time"`  // modify time
	AccountId 				string 			`json:"account_id"`
	TartgetPosition		 	float64       	`json:"target_position"`
	TargetOffset         	float64       	`json:"target_offset"`
}