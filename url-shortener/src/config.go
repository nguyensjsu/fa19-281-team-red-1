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
	BaseURL : "http://54.188.141.200:8080/",
	Backend : "mongo",
	Mongo : MongoConfig {
		Host : "mongodb://admin:cmpe281@10.0.1.144:27017,10.0.1.174:27017,10.0.1.225:27017/admin?replicaSet=cmpe281",
		DB : "shortenURL",
		Collection : "urls",
		Timeout : 10,
	},
}

func GetConfiguration() Configuration {
	return Config
}