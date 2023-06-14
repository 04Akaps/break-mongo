package document

import (
	"break-mongo/types"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func (s *ShardingDocument) AddShardDummy() {
	models := []mongo.WriteModel{}
	option := options.BulkWrite()

	for i := 1; i < 5; i++ {

		data := &types.Orders{}

		data.UserId = strconv.FormatInt(int64(i+1), 10)

		if i == 1 {
			data.Address = "Dummy Address One"
		} else if i == 2 {
			data.Address = "Dummy Address Two"
		} else {
			data.Address = "Dummy Address Three"
		}

		newModel := mongo.NewInsertOneModel().SetDocument(data)

		models = append(models, newModel)
	}

	_, err := s.shardCollection.BulkWrite(context.Background(), models, option)

	if err != nil {
		log.Println(err)
	}
}
