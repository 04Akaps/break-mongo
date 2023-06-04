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
	initDB bool
	DB     *mongo.Database
}

func NewMongoConnect(ctx context.Context, uri string) (*MongoClient, error) {

	mongoDBinit.Do(func() {
		var c *mongo.Client
		var err error

		mongoConn := options.Client().ApplyURI(uri)

		if c, err = mongo.Connect(ctx, mongoConn); err != nil {
			mongoDBInitError = err
			return
		}

		if err = c.Ping(ctx, nil); err != nil {
			mongoDBInitError = err
			return
		}
		client = &MongoClient{Client: c}

	})

	return client, mongoDBInitError

}

func (m *MongoClient) SetMongoDB(dbName string) {
	// 프로젝트 시작 할 떄사용할 DB는 단 하나만 사용 가능하게 구성
	// 여러개 해도 상관은 없음
	if m.initDB {
		panic("DB is Already Setted")
	}

	m.DB = m.Client.Database(dbName)
}
