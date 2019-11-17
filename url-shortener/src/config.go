package main

type Configuration struct {
	BaseURL string
	Backend string
	Mongo	MongoConfig 
}

type MongoConfig struct {
	Host string
	DB	string
	Collection string
	Timeout int
}

var Config = Configuration {
	BaseURL : "http://localhost:8080/",
	Backend : "mongo",
	Mongo : MongoConfig {
		Host : "mongodb://localhost:27017",
		DB : "shortenURL",
		Collection : "urls",
		Timeout : 10,
	},
}

func GetConfiguration() Configuration {
	return Config
}