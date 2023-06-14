package document

import (
	"break-mongo/mongo"
	"context"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	mongo              *mongo.MongoClient
	userCollection     *m.Collection
	userBulkCollection *m.Collection

	fileBucket *gridfs.Bucket

	aggregateCollectionOne *m.Collection
	aggregateCollectionTwo *m.Collection
}

type ShardingDocument struct {
	mongo           *mongo.MongoClient
	shardCollection *m.Collection
}

func NewDocument() *Document {
	ctx := context.Background()
	endPoint := "mongodb+srv://hojin:1111@msggo.wbwdsv8.mongodb.net/?retryWrites=true&w=majority"

	if client, err := mongo.NewMongoConnect(ctx, endPoint); err != nil {
		panic(err)
	} else {
		d := &Document{
			mongo: client,
		}

		d.mongo.DB = d.mongo.Client.Database("break-mongo")

		d.userCollection = d.mongo.DB.Collection("user-collection")
		d.userBulkCollection = d.mongo.DB.Collection("user-bulk-collection")

		d.aggregateCollectionOne = d.mongo.DB.Collection("aggregate-collection-one")
		d.aggregateCollectionTwo = d.mongo.DB.Collection("aggregate-collection-two")

		bucketOpts := options.GridFSBucket().SetName("my-custom-bucket-name")
		if d.fileBucket, err = gridfs.NewBucket(d.mongo.DB, bucketOpts); err != nil {
			panic(err)
		}

		return d
	}

}

func NewShardingDocument() *ShardingDocument {

	ctx := context.Background()
	endPoint := "mongodb://localhost:27017"

	if shardClient, err := mongo.NewMongoConnect(ctx, endPoint); err != nil {
		panic(err)
	} else {

		s := &ShardingDocument{
			mongo: shardClient,
		}

		s.mongo.DB = shardClient.Client.Database("shard-db")
		s.shardCollection = s.mongo.DB.Collection("orders")

		return s
	}

}
