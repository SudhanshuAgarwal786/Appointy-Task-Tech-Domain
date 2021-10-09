package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ParticipantsBusy : Checks if the participants are not RSVP in any other meeting during this time
func UsersBusy(thismeet User) error {
	collection := client.Database("appointytask").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var meet User
	for _, thisperson := range thismeet.Users {
		if thisperson.Password == "Yes" {
			filter := bson.M{
				"users.email": thisperson.Email,
				"users.password":  "Yes",
				"endtime":            bson.M{"$gt": string(time.Now().Format(time.RFC3339))},
			}
			cursor, _ := collection.Find(ctx, filter)
			for cursor.Next(ctx) {
				cursor.Decode(&meet)
				if (thismeet.Starttime >= meet.Starttime && thismeet.Starttime <= meet.Endtime) ||
					(thismeet.Endtime >= meet.Starttime && thismeet.Endtime <= meet.Endtime) {
					returnerror := "Error 400: User " + thisperson.Name + " Password Clash"
					return errors.New(returnerror)
				}
			}
		}
	}
	return nil
}

//CreateMeeting : Adds another meeting to the database
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meet User
	_ = json.NewDecoder(request.Body).Decode(&meet)
	meet.def()
	if meet.Starttime < meet.Creationtime {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Meeting cannot start in the past" }`))
		return
	}
	if meet.Starttime > meet.Endtime {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid time" }`))
		return
	}
	lock.Lock()
	defer lock.Unlock()
	err := UsersBusy(meet)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	collection := client.Database("appointy").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, meet)
	meet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(meet)
	fmt.Println(meet)
}
