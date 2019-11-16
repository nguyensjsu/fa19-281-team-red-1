package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type userinfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type userinfoDB struct {
	Username string `json:"Username"`
	Password []byte `json:"Password"`
}

var c *mgo.Collection

func main() {
	c = initDB()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", root).Methods("Get")
	router.HandleFunc("/hello", hello).Methods("Get")
	router.HandleFunc("/headers", headers).Methods("Get")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/signin", signin).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() *mgo.Collection {
	url := "localhost:27017"
	database := "user_auth"
	collection := "userinfo"

	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("[initDB] " + err.Error())
	}

	return session.DB(database).C(collection)
}

func signup(w http.ResponseWriter, req *http.Request) {
	log.Println("In signup")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	var user userinfo
	json.Unmarshal(reqBody, &user)
	log.Println("Username: " + user.Username + " Password: " + user.Password)

	if len(user.Username) == 0 || len(user.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Signup Error")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	userDB := userinfoDB{user.Username, hashedPassword}

	c.Insert(&userDB)

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "User Created")
}

func signin(w http.ResponseWriter, req *http.Request) {
	log.Println("In signin")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	var user userinfo
	json.Unmarshal(reqBody, &user)
	log.Println("Username: " + user.Username + " Password: " + user.Password)

	var result userinfoDB
	err = c.Find(bson.M{"username": user.Username}).One(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	log.Println("Username: " + result.Username + " Password: " + string(result.Password))

	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	fmt.Fprintf(w, "[In Signin] Username: "+result.Username)
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
