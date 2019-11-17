package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"golang.org/x/crypto/bcrypt"
)

type userinfo struct {
	Username string `json:"Username" bson:"Username"`
	Password string `json:"Password" bson:"Password"`
}

type userinfoDB struct {
	Username string   `json:"Username" bson:"Username"`
	Password []byte   `json:"Password" bson:"Password"`
	History  []string `json:"History" bson:"History"`
}

// type userinfohistory struct {
// 	Id       bson.ObjectId `json:"_id" bson:"_id"`
// 	Username string        `json:"Url" bson:"Url"`
// 	Password string        `json:"Password" bson:"Password"`
// 	History  []string      `json:"History" bson:"History"`
// }

type requestStruct struct {
	Url      string
	Username string
}

var c *mgo.Collection
var formatter *render.Render

func main() {
	formatter = render.New(render.Options{
		IndentJSON: true,
	})

	c = initDB()
	n := negroni.Classic()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/headers", headers).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/signin", signin).Methods("POST")
	router.HandleFunc("/url", urlHandler).Methods("POST")
	router.HandleFunc("/history", historyHandler).Methods("GET")
	n.UseHandler(router)
	// http.ListenAndServe(":8080", n)
	n.Run(":8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() *mgo.Collection {
	// Local
	// url := "localhost:27017"
	// Docker-compose
	url := "mongodb:27017"

	database := "user_auth"
	collection := "userinfo"

	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("[initDB] " + err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(database).C(collection)
}

// Create history for a user
func urlHandler(w http.ResponseWriter, req *http.Request) {
	var request requestStruct
	_ = json.NewDecoder(req.Body).Decode(&request)
	fmt.Println("URL in request: ", request.Url)
	fmt.Println("Username in request: ", request.Username)
	query := bson.M{"Username": request.Username}
	var user userinfoDB
	err := c.Find(query).One(&user)
	if err != nil {
		formatter.JSON(w, http.StatusNotFound, "User not found")
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "User not found")
		return
		// log.Fatal(err)
	}
	fmt.Println(user.Username)
	// fmt.Println(user.History)
	user.History = append(user.History, request.Url)
	fmt.Println(user.History)
	change := bson.M{"$set": bson.M{"History": user.History}}
	err = c.Update(query, change)
	if err != nil {
		formatter.JSON(w, http.StatusInternalServerError, "Could not update database")
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintf(w, "Could not update database")
		return
	}
	formatter.JSON(w, http.StatusCreated, "Successfully added to history")
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "Successfully added to history")

}

func historyHandler(w http.ResponseWriter, req *http.Request) {

	username, ok := req.URL.Query()["Username"]
    
    if !ok || len(username[0]) < 1 {
		formatter.JSON(w, http.StatusBadRequest, "Url Param 'Username' is missing")
        return
	}
	
	query := bson.M{"Username": username[0]}
	var result userinfoDB
	err := c.Find(query).One(&result)
	if err != nil {
		formatter.JSON(w, http.StatusNotFound, "User not found")
		return
	}
	fmt.Println("Results All: ", result)
	formatter.JSON(w, http.StatusOK, result)
}

func signup(w http.ResponseWriter, req *http.Request) {
	log.Println("In signup")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		formatter.JSON(w, http.StatusBadRequest, "Signup Error: "+err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	var user userinfo
	json.Unmarshal(reqBody, &user)
	log.Println("Username: " + user.Username + " Password: " + user.Password)

	if len(user.Username) == 0 || len(user.Password) == 0 {
		formatter.JSON(w, http.StatusBadRequest, "Signup Error: "+err.Error())
		// w.WriteHeader(http.StatusBadRequest)

		// fmt.Fprintf(w, "Signup Error")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	userDB := userinfoDB{user.Username, hashedPassword, make([]string, 0)}
	c.Insert(&userDB)
	formatter.JSON(w, http.StatusCreated, "User Created")

	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "User Created")
}

func signin(w http.ResponseWriter, req *http.Request) {
	log.Println("In signin")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		formatter.JSON(w, http.StatusBadRequest, "Signin Error: "+err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	var user userinfo
	json.Unmarshal(reqBody, &user)
	log.Println("Username: " + user.Username + " Password: " + user.Password)

	var result userinfoDB
	err = c.Find(bson.M{"username": user.Username}).One(&result)
	if err != nil {
		formatter.JSON(w, http.StatusBadRequest, "Signin Error: "+err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprintf(w, "Signup Error: "+err.Error())
		return
	}

	log.Println("Username: " + result.Username + " Password: " + string(result.Password))

	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		// w.WriteHeader(http.StatusUnauthorized)
		formatter.JSON(w, http.StatusUnauthorized, "Unauthorized")
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
