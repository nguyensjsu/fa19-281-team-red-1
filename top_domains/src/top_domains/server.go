package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
// var mongodb_server = "mongodb"

// Local
// var mongodb_server = "localhost:27017"
// var mongodb_database = "top_domains"
// var mongodb_collection = "top_domains"

// Docker-compose
// var mongodb_server = "mongodb:27017"
// var mongodb_database = "top_domains"
// var mongodb_collection = "top_domains"

// AWS EC2 instance
var mongodb_server = "mongodb://admin:test@10.0.1.87:27017"
var mongodb_database = "top_domains"
var mongodb_collection = "top_domains"

// Get TOP LIMIT of most popular hit websites
// const Limit = 5
var limit = 5

// RabbitMQ Config
// var rabbitmq_server = "rabbitmq"
// var rabbitmq_port = "5672"
// var rabbitmq_queue = "gumball"
// var rabbitmq_user = "guest"
// var rabbitmq_pass = "guest"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/domains", domainHandler(formatter)).Methods("GET")
	mx.HandleFunc("/url", urlHandler(formatter)).Methods("POST")
	mx.HandleFunc("/top", topUrlHandler(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Top Domains API version 1.0 alive!"})
	}
}

// API Update Old Domain or Create New Domain
func urlHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var inputUrl urlStruct
		// var d domainMap
		_ = json.NewDecoder(req.Body).Decode(&inputUrl)
		fmt.Println("URL in request: ", inputUrl.Url)
		u, err := url.Parse(inputUrl.Url)
		if err != nil {
			panic(err)
		}
		// fmt.Println(u.Scheme)
		// fmt.Println(u.Host)
		components := strings.Split(u.Host, ".")
		domain, _ := strings.ToLower(components[0]), components[1]
		fmt.Println("Domain is ", domain)

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		query := bson.M{"Domain": domain}
		count, err := c.Find(query).Count()
		var d domainMap
		if count > 0 {
			fmt.Println("Found an instance of this domain, UPDATE COUNTER")
			err = c.Find(query).One(&d)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(d)
			fmt.Println("Counter result: ", d.Counter)
			d.Counter = d.Counter + 1
			d.Url = append(d.Url, inputUrl.Url)
			fmt.Println("NEW Counter result: ", d.Counter)
			change := bson.M{"$set": bson.M{"Counter": d.Counter, "Url": d.Url}}
			err = c.Update(query, change)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Updated domain and counter: ", domain, d.Counter)
			formatter.JSON(w, http.StatusPartialContent, d)
		} else {
			fmt.Println("Did not find in DB, going to create a new instance")
			d = domainMap{
				Id:      bson.NewObjectId(),
				Url:     []string{inputUrl.Url},
				Counter: 1,
				Domain:  domain,
			}
			err := c.Insert(d)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created new document", domain)
			formatter.JSON(w, http.StatusCreated, d)

		}
	}
}

// GET top LIMIT Url Handler
func topUrlHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var d []domainMap
		err = c.Find(nil).Sort("-Counter").Select(bson.M{"_id": 1, "Counter": 1, "Domain": 1}).Limit(limit).All(&d)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Results: ", d)
		formatter.JSON(w, http.StatusOK, d)
	}
}

// GET all objects from collection
func domainHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			fmt.Println("Failed to establish connection to Mongo server:", err)
			// panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		// var result bson.M
		result := []domainMap{}
		err = c.Find(nil).All(&result)
		if err != nil {
			fmt.Println(err) // prints 'document is nil'
		}
		// if err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("Results All: ", result)
		// fmt.Println("Gumball Machine:", result)
		formatter.JSON(w, http.StatusOK, result)
	}

}

/*

  	-- Gumball MongoDB Collection (Create Document) --

    db.gumball.insert(
	    {
	      Id: 1,
	      CountGumballs: NumberInt(202),
	      ModelNumber: 'M102988',
	      SerialNumber: '1234998871109'
	    }
	) ;

    -- Gumball MongoDB Collection - Find Gumball Document --

    db.gumball.find( { Id: 1 } ) ;

    {
        "_id" : ObjectId("54741c01fa0bd1f1cdf71312"),
        "Id" : 1,
        "CountGumballs" : 202,
        "ModelNumber" : "M102988",
        "SerialNumber" : "1234998871109"
    }

    -- Gumball MongoDB Collection - Update Gumball Document --

    db.gumball.update(
        { Dd: 1 },
        { $set : { CountGumballs : NumberInt(10) } },
        { multi : false }
    )

    -- Gumball Delete Documents

    db.gumball.remove({})

*/
