package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World from handler")
}

func HomeHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World from Home")
}

func CreateUserHandler(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	response, err := user.toJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
