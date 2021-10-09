package main

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var lock sync.Mutex

//Defaultskip : Stores the value of the default offset
var Defaultskip = int64(0)

//Defaultlimit : Stores the value of the default limit
var Defaultlimit = int64(10)

var skip = Defaultskip
var limit = Defaultlimit

//participant : Stores the details of a participant
type user struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
}

//Meeting : Stores the details of a meeting
type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Users		 []user             `json:"users,omitempty" bson:"users,omitempty"`
	Starttime    string             `json:"starttime,omitempty" bson:"starttime,omitempty"`
	Endtime      string             `json:"endtime,omitempty" bson:"endtime,omitempty"`
	Creationtime string             `json:"creationtime,omitempty" bson:"creationtime,omitempty"`
}

func (person *user) cons() {
	if person.Password == "" {
		person.Password = "Not Answered"
	}
	if person.Email == "" {
		person.Email = "defaultmail@email.com"
	}
	if person.Name == "" {
		person.Name = person.Email
	}
}

func (obj *User) def() {
	if obj.Title == "" {
		obj.Title = "Untitled Meeting"
	}
	if obj.Starttime == "" {
		obj.Starttime = string(time.Now().Format(time.RFC3339))
	}
	if obj.Endtime == "" {
		obj.Endtime = string(time.Now().Local().Add(time.Hour * time.Duration(1)).Format(time.RFC3339))
	}
	if obj.Creationtime == "" {
		obj.Creationtime = string(time.Now().Format(time.RFC3339))
	}
	for i := range obj.Users {
		obj.Users[i].cons()
	}
}
