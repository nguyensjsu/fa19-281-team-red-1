package main

type Configuration struct {
	BaseURL string
	Backend string
	TopServiceURL string
	UserServiceURL string
	Mongo	MongoConfig 
}

type MongoConfig struct {
	Host string
	DB	string
	Collection string
	Timeout int
}

var Config = Configuration {
	BaseURL : "http://ec2-18-236-198-15.us-west-2.compute.amazonaws.com:8000/unshorten/",
	Backend : "mongo",
	TopServiceURL : "http://go-api-project-2106290823.us-west-2.elb.amazonaws.com/url",
	UserServiceURL : "http://url-elb-1512171399.us-west-2.elb.amazonaws.com/url",
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