package document

import (
	"break-mongo/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (d *Document) AggregateDummy() {
	models := []mongo.WriteModel{}
	option := options.BulkWrite()

	firstModel := types.AggregateDummyOne{
		Name:         "FaceBook",
		CategoryCode: "social",
		Description:  "first Model",
		FoundYear:    2004,
		FundingRounds: []types.FundingRoundsART{
			{
				RoundCode:    "a",
				RaisedAmount: 150,
				FundedYear:   2007,
				Investments: []types.Investment{
					{
						Company:    "Google",
						Person:     false,
						FundAmount: 80,
					},
					{
						Company:    "MS",
						Person:     false,
						FundAmount: 70,
					},
				},
			},
		},
	}
	secondModel := types.AggregateDummyOne{
		Name:         "MS",
		CategoryCode: "IT",
		Description:  "Second Model",
		FoundYear:    2000,
		FundingRounds: []types.FundingRoundsART{
			{
				RoundCode:    "b",
				RaisedAmount: 300,
				FundedYear:   2001,
				Investments: []types.Investment{
					{
						Company:    "",
						Person:     true,
						FundAmount: 300,
					},
				},
			},
		},
	}
	thirdModel := types.AggregateDummyOne{
		Name:         "Test",
		CategoryCode: "social",
		Description:  "Third Model",
		FoundYear:    2023,
		FundingRounds: []types.FundingRoundsART{
			{
				RoundCode:    "a",
				RaisedAmount: 1000,
				FundedYear:   2023,
				Investments: []types.Investment{
					{
						Company:    "",
						Person:     true,
						FundAmount: 300,
					},
				},
			},
			{
				RoundCode:    "b",
				RaisedAmount: 50,
				FundedYear:   2023,
				Investments: []types.Investment{
					{
						Company:    "unknown",
						Person:     false,
						FundAmount: 50,
					},
				},
			},
		},
	}

	models = append(models, mongo.NewInsertOneModel().SetDocument(firstModel))
	models = append(models, mongo.NewInsertOneModel().SetDocument(secondModel))
	models = append(models, mongo.NewInsertOneModel().SetDocument(thirdModel))

	_, err := d.aggregateCollectionOne.BulkWrite(context.Background(), models, option)

	if err != nil {
		log.Println(err)
	}

	modelsTwo := []mongo.WriteModel{}

	firstModelTwo := types.AggregateDummyTwo{
		Name:    "FaceBook",
		Age:     30,
		Address: "Seoul",
	}

	secondModelTwo := types.AggregateDummyTwo{
		Name:    "MS",
		Age:     50,
		Address: "Home",
	}

	modelsTwo = append(modelsTwo, mongo.NewInsertOneModel().SetDocument(firstModelTwo))
	modelsTwo = append(modelsTwo, mongo.NewInsertOneModel().SetDocument(secondModelTwo))
	_, err = d.aggregateCollectionTwo.BulkWrite(context.Background(), modelsTwo, option)

	if err != nil {
		log.Println(err)
	}

}

func (d *Document) FindAllAggregateOne() (*[]types.AggregateDummyOne, error) {

	filter := bson.M{}

	opts := []*options.FindOptions{}
	result := []types.AggregateDummyOne{}

	cursor, err := d.aggregateCollectionOne.Find(context.Background(), filter, opts...)
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

func (d *Document) FindAllAggregateTwo() (*[]types.AggregateDummyTwo, error) {
	filter := bson.M{}

	opts := []*options.FindOptions{}
	result := []types.AggregateDummyTwo{}

	cursor, err := d.aggregateCollectionTwo.Find(context.Background(), filter, opts...)
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

func (d *Document) FindByNameAggregate(name string) (*[]types.AggregateDummyOne, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"name": name,
			},
		},
		{
			"$lookup": bson.M{
				"from":         d.aggregateCollectionOne.Name(),
				"localField":   "name",
				"foreignField": "name",
				"as":           "joinData",
			},
		},
		{
			"$unwind": "$joinData",
		},
		{
			"$project": bson.M{
				"name":          "$name",
				"categoryCode":  "$joinData.categoryCode",
				"description":   "$joinData.description",
				"foundYear":     "$joinData.foundYear",
				"fundingRounds": "$joinData.fundingRounds",
			},
		},
		//{
		//	"$match": bson.M{
		//		"foundYear": bson.M{
		//			"$lte": bson.M{
		//				"foundYear": 3000,
		//			},
		//		},
		//	},
		//},
	}

	ctx := context.TODO()
	cursor, err := d.aggregateCollectionTwo.Aggregate(ctx, pipeline)

	if err != nil {

		fmt.Println("1", err)
	}

	result := []types.AggregateDummyOne{}

	err = cursor.All(ctx, &result)

	if err != nil {
		fmt.Println(err)
	}

	return &result, nil

}
