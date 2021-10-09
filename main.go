package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Sudhanshu's server application running")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/user/", GetUsers)
	http.HandleFunc("/posts/", GetUserwithID)
	http.HandleFunc("/posts/users", GetUserwithTime)
	http.ListenAndServe(":5658", nil)
}
