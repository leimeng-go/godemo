package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

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
