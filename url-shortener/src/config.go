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
	BaseURL : "http://54.203.79.184:8080/",
	Backend : "mongo",
	Mongo : MongoConfig {
		Host : "mongodb://admin:cmpe281@primary:27017,secondary1:27017,secondary2:27017/admin?replicaSet=cmpe281",
		//Host : "mongodb://root:0717@localhost:27017",
		DB : "shortenURL",
		Collection : "urls",
		Timeout : 10,
	},
}

func GetConfiguration() Configuration {
	return Config
}