package types

type BulkUser struct {
	Name       string `json:"name" bson:"name"`
	SecondName string `json:"secondName" bson:"secondName"`
	Age        int64  `json:"age" bson:"age"`
}
