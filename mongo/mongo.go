package mongo

import (
	"context"
	"time"
     "log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewMongoClient(url string){
   ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
   defer cancel()
   client,err:=mongo.Connect(ctx,options.Client().ApplyURI("mongodb://root:181205@mongo-1:27011,mongo-2:27012,mongo-3:27013/?replicaSet=rs&authSource=admin"))
   if err!=nil{
	 log.Error(err.Error())
   }
   
}