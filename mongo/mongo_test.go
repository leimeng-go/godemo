package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)
func newDefaultClient()*mongo.Client{
	url:="mongodb://root:181205@mongo-1:27011,mongo-2:27012,mongo-3:27013/?replicaSet=rs&authSource=admin"
	client:=NewMongoClient(url)
	ctx:=context.Background()
	err :=client.Ping(ctx,readpref.Primary())
	if err!=nil{
		log.Fatal(err.Error())
		return nil
	}
	return client
}
func TestNewMongoClient(t *testing.T) {
	url:="mongodb://root:181205@mongo-1:27011,mongo-2:27012,mongo-3:27013/?replicaSet=rs&authSource=admin"
	client:=NewMongoClient(url)
	ctx:=context.Background()
	err :=client.Ping(ctx,readpref.Primary())
	if err!=nil{
		t.Log(err.Error())
		return
	}
    t.Log("mongo connect success")
}

func TestReduce(t *testing.T){
	par:=bson.D{
		{"mapreduce", "xw_landlords"},
		{"map", " function() { emit( this.nickname , 1 ); }"},
		{"reduce", "function(key, arr) { return Array.sum(arr); }"},
		{"out", "mr3"},
		{"query", bson.D{{"status", "SUCCESS"}}},
	}
	client:=newDefaultClient()
	ctx:=context.Background()
	result,err :=client.Database("zxw").RunCommand(ctx,par).DecodeBytes()
	if err!=nil{
		t.Log(err.Error())
	}
	t.Logf("%+v",result)
}