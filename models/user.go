package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID     bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthId string        `json:"authId" bson:"authId"`
	Name   string        `json:"name" bson:"name"`
	Role   []Role        `json:"roles" bson:"role"`
}

type Role struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Permissions []Permission  `jsoin:"permissions" bson:"permissions"`
}

type Permission struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}
