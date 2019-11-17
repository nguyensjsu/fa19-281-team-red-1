package main

import (
	"context"
	"fmt"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


type Store struct {
	c *mongo.Client
}

type Result struct {
	Shortened string
	Origin string
}

func NewDB() (*Store, error) {
	fmt.Println("[NewDB] Mongo host: ", conf.Mongo.Host)
	clientOptions := options.Client().ApplyURI(conf.Mongo.Host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("[NewDB] Can NOT connect to mongo db: ", err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("[NewDB] Ping mongo failed")
		return nil, err
	}
	fmt.Println("[NewDB] Connected to mongo db")
	ret := &Store{c:client}
	return ret, nil
}

func (s *Store) CreateEntry(short string, origin string) error {
	collection := s.c.Database(conf.Mongo.DB).Collection(conf.Mongo.Collection)
	filter := bson.M{"shortened": short}
	var result Result
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		fmt.Println("[CreateEntry] Shortened URL Exisited")
		return errors.New("Shortened URL Exisited")
	}
	_, err = collection.InsertOne(context.TODO(), bson.M{"shortened": short, "origin": origin})
	if err != nil {
		fmt.Println("[CreateEntry] Insert into DB failed: ", err)
		return errors.New("Insert into DB failed")
	}
	return nil
}

func (s *Store) GetEntry(shortUrl string) (string, error) {
	collection := s.c.Database(conf.Mongo.DB).Collection(conf.Mongo.Collection)
	filter := bson.M{"shortened": shortUrl}
	//filter := bson.M{}
	result := Result{Shortened : "xxx", Origin: "xxx",}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("[GetEntry] Find shortened URL failed: ", err)
		return "", errors.New("Find shortened URL failed")
	}
	fmt.Println("[GetEntry] Get resutl from mongo: ", result)
	return result.Origin, nil
}

func (s *Store) Close() error {
	return  s.c.Disconnect(context.TODO())
}