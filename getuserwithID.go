package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckMeetingwithID : Checks the meeting with the provided id
func CheckUserwithID(id primitive.ObjectID) (User, error) {
	var meet User
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, User{ID: id}).Decode(&meet)
	if meet.ID != id {
		err = errors.New("error 400: ID not present")
	}
	return meet, err
}

//GetMeetingwithID : Gives the meeting with the provided id
func GetUserwithID(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println(path.Base(request.URL.Path))
	id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
	meetingwithID, err := CheckUserwithID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meetingwithID)
	
}
