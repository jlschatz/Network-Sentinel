package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	Id  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IP  string             `json:"ip" bson:"ip,omitempty"`
	MAC string             `json:"mac" bson:"mac,omitempty"`
}
