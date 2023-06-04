package document

import (
	"break-mongo/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (d *Document) InsertDummyUserData(req *types.User) {
	filter := bson.M{
		"name": req.Name,
		"age":  req.Age,
	}

	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": bson.M{
			"name": req.Name,
			"age":  req.Age,
		},
	}

	if _, err := d.userCollection.UpdateOne(context.Background(), filter, update, opts); err != nil {
		panic(err)
	}
}

func (d *Document) InsertBulkDummyUserData(req *types.User) {

}

func (d *Document) UpdateUserData(req *types.User) error {
	filter := bson.M{
		"name": req.Name,
		"age":  req.Age,
	}

	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": bson.M{
			"name": req.Name,
		},

		"$inc": bson.M{
			"age": 100,
		},
	}

	if _, err := d.userCollection.UpdateOne(context.Background(), filter, update, opts); err != nil {
		return err
	}
	return nil
}

func (d *Document) FindUserData() (*[]types.User, error) {
	// Find All
	filter := bson.M{}
	opts := []*options.FindOptions{}
	result := []types.User{}

	cursor, err := d.userCollection.Find(context.Background(), filter, opts...)
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

func (d *Document) FindOneUserData(req *types.User) (*types.User, error) {
	res := types.User{}

	filter := bson.M{
		"name": req.Name,
		"age":  req.Age,
	}
	opts := []*options.FindOneOptions{}

	if err := d.userCollection.FindOne(context.Background(), filter, opts...).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (d *Document) FindUserDataByAgeSort() (*[]types.User, error) {
	filter := bson.M{}
	opts := []*options.FindOptions{
		options.Find().SetSort(bson.M{"age": -1}),
		options.Find().SetProjection(bson.M{"name": 1, "age": 1}),
		options.Find().SetLimit(10),
	}

	result := []types.User{}

	cursor, err := d.userCollection.Find(context.Background(), filter, opts...)
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
