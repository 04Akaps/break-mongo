package types

type Orders struct {
	UserId  string `json:"user_id" bson:"user_id"`
	Address string `json:"address" bson:"address"`
}
