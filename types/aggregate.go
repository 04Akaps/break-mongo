package types

type AggregateDummyTwo struct {
	Name    string `json:"name" bson:"name"`
	Age     int64  `json:"age" bson:"age"`
	Address string `json:"address" bson:"address"`
}

type AggregateDummyOne struct {
	Name          string             `json:"name" bson:"name"`
	CategoryCode  string             `json:"categoryCode" bson:"categoryCode"`
	Description   string             `json:"description" bson:"description"`
	FoundYear     int64              `json:"foundYear" bson:"foundYear"`
	FundingRounds []FundingRoundsART `json:"fundingRounds" bson:"fundingRounds"`
}

type FundingRoundsART struct {
	RoundCode    string       `json:"round_code" bson:"round_code"`
	RaisedAmount int64        `json:"raised_amount" bson:"raised_amount"`
	FundedYear   int64        `json:"funded_year" bson:"funded_year"`
	Investments  []Investment `json:"investments" bson:"investments"`
}

type Investment struct {
	Company    string `json:"company" bson:"company"`
	Person     bool   `json:"person" bson:"person"`
	FundAmount int64  `json:"fundAmount" bson:"fundAmount"`
}
