package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func NewMongoClient(url string)*mongo.Client{
   ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
   defer cancel()
   client,err:=mongo.Connect(ctx,options.Client().ApplyURI(url))
   if err!=nil{
	 log.Fatal(err.Error())
	 return nil
   }
	return client
}