package document

import (
	"break-mongo/mongo"
	"context"
	m "go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	mongo              *mongo.MongoClient
	userCollection     *m.Collection
	userBulkCollection *m.Collection
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

		d.mongo.SetMongoDB("break-mongo")

		d.userCollection = d.mongo.DB.Collection("user-collection")
		d.userBulkCollection = d.mongo.DB.Collection("user-bulk-collection")
		return d
	}

}
