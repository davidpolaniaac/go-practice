package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) toJson() ([]byte, error) {
	return json.Marshal(u)
}
