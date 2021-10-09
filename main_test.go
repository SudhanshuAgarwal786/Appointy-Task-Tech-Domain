package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMainCreatemeet(t *testing.T) {
	var message User
	var part user
	part.Name = "Sudhanshu Agarwal"
	part.Email = "sudhanshusanjay8@gmail.com"
	part.Password = "Not found"
	message.Title = "Title"
	message.Users = append(message.Users, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"
	bytesRepresentation, _ := json.Marshal(message)
	resp, err := http.Post("http://localhost:12345/meetings", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}
func TestMaingetmeet(t *testing.T) {
	resp, err := http.Get("http://localhost:12345/user/?id=5f4dcb738b246dc74d8ecd44")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}
func TestMaingetparticipants(t *testing.T) {
	resp, err := http.Get("http://localhost:12345/posts/?participant=abcd@gmail.com")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}
func TestMaingetmeetintime(t *testing.T) {
	resp, err := http.Get("http://localhost:12345/meetings?start=2019-09-01T13:30:10+05:30&end=2021-09-01T14:30:10+05:30")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}

func BenchmarkMaingetmeet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:12345/meeting/5f4dcc74fa1a4b2011daf69a")
	}
}

func BenchmarkMaingetparticipant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:12345/posts/?user=sudhanshusanjay8@gmail.com")
	}
}

func BenchmarkMaingettime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:12345/meetings?start=2019-09-01T13:30:10+05:30&end=2021-09-01T14:30:10+05:30")
	}
}

func BenchmarkMaingetpost(b *testing.B) {
	var message User
	var part user
	part.Name = "xyz"
	part.Email = "xyz.com"
	part.Password = "No"
	message.Title = "Title"
	message.Users = append(message.Users, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error")
	}

	for n := 0; n < b.N; n++ {
		resp, err := http.Post("http://localhost:12345/meetings", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			b.Error("Fail")
		}
		if resp == nil {
			b.Error("NO response")
		}
	}
}

func BenchmarkParticipantsBusy(b *testing.B) {
	var message User
	var part user
	part.Name = "xyz"
	part.Email = "xyz@gmail.com"
	part.Password = "No"
	message.Title = "Title"
	message.Users = append(message.Users, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)
	for n := 0; n < b.N; n++ {
		UsersBusy(message)
	}
}

func BenchmarkCheckUser(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)
	for n := 0; n < b.N; n++ {
		_ = CheckUser("xyz@gmail.com")
	}
}
