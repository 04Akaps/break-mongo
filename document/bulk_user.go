package document

import (
	"break-mongo/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (d *Document) InsertBulkDummyUserData() {
	models := []mongo.WriteModel{}
	option := options.BulkWrite()

	for i := 1; i < 100000; i++ {
		var name string
		var secnodName string

		if i < 100 {
			name = "A"
			secnodName = "AA"
		} else if i < 500 {
			name = "B"
			secnodName = "BB"
		} else {
			name = "C"
			secnodName = "CC"
		}

		newModel := mongo.NewInsertOneModel().SetDocument(types.BulkUser{
			Name:       name,
			SecondName: secnodName,
			Age:        int64(i + 1),
		})

		models = append(models, newModel)
	}

	_, err := d.userBulkCollection.BulkWrite(context.Background(), models, option)

	if err != nil {
		log.Println(err)
	}

}

func (d *Document) FindAllBulkUser(f bool) (*[]types.BulkUser, error) {

	filter := bson.M{}

	if f {
		filter["name"] = "A"
	}

	opts := []*options.FindOptions{}
	result := []types.BulkUser{}

	cursor, err := d.userBulkCollection.Find(context.Background(), filter, opts...)
	defer cursor.Close(context.Background())

	if err != nil {
		log.Println("Find Error ", err.Error())
		return nil, err
	}

	if err = cursor.All(context.Background(), &result); err != nil {
		log.Println("cursor All Error : ", err)
		return nil, err
	}

	return &result, nil
}

func (d *Document) FindAllBulkUserBySort() (*[]types.BulkUser, error) {

	filter := bson.M{
		"age": bson.M{
			"$gte": 21,
			"$lte": 30,
		},
	}

	opts := options.Find()
	// Name으로 정렬이 일단 되어 있기 떄문에 그렇게 많은 영향을 주지는 않는다.
	opts.SetSort(bson.D{{Key: "name", Value: -1}})

	result := []types.BulkUser{}

	cursor, err := d.userBulkCollection.Find(context.Background(), filter, opts)
	defer cursor.Close(context.Background())

	if err != nil {
		log.Println("Find Error ", err.Error())
		return nil, err
	}

	if err = cursor.All(context.Background(), &result); err != nil {
		log.Println("cursor All Error : ", err)
		return nil, err
	}

	return &result, nil
}

func (d *Document) AddIndexByName() {
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
	}

	if _, err := d.userBulkCollection.Indexes().CreateOne(context.Background(), indexModel); err != nil {
		log.Println(err)
	}
}

func (d *Document) AddDoubleIndex() {
	// 만약 순서를 age, name으로 두면 FindAllBulkUserBySort 함수에서 효율적으로 동작

	// 하지만 반대로 name, age로 두면 name을 기준으로 age를 정렬하기 떄문에
	// FindAllBulkUserBySort에서 비효율적으로 동작을 하게 된다.
	// -> 왜냐하면 FindAllBulkUserBySort 에서는 age기준으로 name을 가져오기 떄문에
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"age", 1},
			{"name", 1},
		},
	}

	_, err := d.userBulkCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
}
