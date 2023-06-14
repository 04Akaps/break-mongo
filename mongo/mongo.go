package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	mongoDBinit      sync.Once
	mongoDBInitError error

	client *MongoClient
)

type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoConnect(ctx context.Context, uri string) (*MongoClient, error) {

	var c *mongo.Client
	var err error

	mongoConn := options.Client().ApplyURI(uri)

	if c, err = mongo.Connect(ctx, mongoConn); err != nil {
		return nil, err
	}

	if err = c.Ping(ctx, nil); err != nil {
		return nil, err
	}

	client = &MongoClient{Client: c}

	return client, mongoDBInitError

}
