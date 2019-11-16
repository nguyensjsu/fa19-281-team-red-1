package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", root).Methods("Get")
	router.HandleFunc("/hello", hello).Methods("Get")
	router.HandleFunc("/headers", headers).Methods("Get")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/signin", signin).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func signup(w http.ResponseWriter, req *http.Request) {
	log.Println("In signup")
	fmt.Fprintf(w, "In Signup")
}

func signin(w http.ResponseWriter, req *http.Request) {
	log.Println("In signin")
	fmt.Fprintf(w, "In Signin")
}

func root(w http.ResponseWriter, req *http.Request) {
	log.Println("In root")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
