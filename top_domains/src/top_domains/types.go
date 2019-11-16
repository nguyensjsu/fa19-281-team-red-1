package main

import "gopkg.in/mgo.v2/bson"

type domainMap struct {
	Id      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Url     []string      `json:"Url,omitempty" bson:"Url,omitempty"`
	Counter int           `json:"Counter,omitempty" bson:"Counter,omitempty"`
	Domain  string        `json:"Domain,omitempty" bson:"Domain,omitempty"`
}

type urlStruct struct {
	Url string
}
