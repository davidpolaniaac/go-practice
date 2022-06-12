package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, isPath := r.rules[path]
	handler, isMethod := r.rules[path][method]
	return handler, isPath, isMethod
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, path, method := r.FindHandler(request.URL.Path, request.Method)
	if !path {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Web site not Found 404")
		return
	}

	if !method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed")
		return
	}
	handler(w, request)
}
