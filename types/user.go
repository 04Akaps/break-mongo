package types

type User struct {
	Name string `json:"name" bson:"name" binding:"required"`
	Age  int64  `json:"age" bson:"age" binding:"required"`
}
