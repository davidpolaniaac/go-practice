package main

func main() {
	server := NewServer(":8080")
	server.Handle("GET", "/", RootHandler)
	server.Handle("POST", "/user", CreateUserHandler)
	server.Handle("GET", "/api", server.AddMiddleware(HomeHandler, Logging(), CheckAuth()))
	server.Listen()
}
