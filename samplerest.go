package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", handler)
	route.HandleFunc("/users", getUsers)
	fmt.Println("Starting the server")
	http.ListenAndServe(":8080", route)
}

func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Landed..!")
	fmt.Fprintf(resp, "Welcome to the HomePage!")
}

func getUsers(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Landed for Get users..!")
}
