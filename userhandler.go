package main

import "net/http"


func UserHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CreateUser(response, request)
	} else if request.Method == "GET" {
		GetUserwithTime(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
