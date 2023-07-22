package document

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (m *Document) SearchTest() {

	searchStage := bson.D{
		{"$search", bson.M{
			"index": "sampleSearchINdex",
			//"text": bson.D{
			//	{"query", "aabbc"},
			//	{"path", "address"},
			//},
			//"returnStoredSource": true,
			//"compound": bson.D{
			//
			//	//{"filter", bson.A{
			//	//	bson.D{{"text", bson.D{{"query", "aabbc"}, {"path", "address"}}}},
			//	//}},
			//	//{"should", bson.A{
			//	//	bson.D{{"text", bson.D{{"query", "tetset"}, {"path", "description"}}}},
			//	//	bson.D{{"text", bson.D{{"query", "aabbc"}, {"path", "address"}}}},
			//	//}},
			//	//{"must", bson.A{
			//	//	//bson.D{{"text", bson.D{{"query", "aabbc"}, {"path", "address"}}}},
			//	//	//bson.D{{"wildcard", bson.D{{"path", "address"}, {"query", "*ㅋ*"}, {"allowAnalyzedField", true}}}},
			//	//	//bson.D{{"exists", bson.D{{"path", "userOffer"}}}},
			//	//	//bson.D{{"range", bson.D{{"path", "tid"}, {"gte", 3}, {"lte", 4}}}},
			//	//	//bson.D{{"queryString", bson.D{
			//	//	//	{"defaultPath", "address"},
			//	//	//	{"query", "(ㅋㅋㅋ OR aabbc) AND description : tetset"},
			//	//	//}}},
			//	//}},
			//	//{"mustNot", bson.A{
			//	//	bson.D{{"wildcard", bson.D{{"path", "description"}, {"query", "*"}, {"allowAnalyzedField", true}}}},
			//	//	//bson.D{{"text", bson.D{{"path", "description"}, {"query", "tetset"}}}},
			//	//}},
			//},
		},
		},
	}

	// Search는 string형태만 가능
	// int타입에 대해서는 따로 match파이프라인을 조회

	// should는 or조건
	// must는 and조건

	// filter, must는 차이가 없음 -> 있다면 그냥 score에 대한 영향 정도
	// mustNot은 != 조건과 동일
	// -> 와일드 카드를 사용하기 위해서는 wildcard를 사용해야 한다.
	// -> must에도 적용 가능

	// null 체크는 exists로 처리
	// embeddedDocument -> 도큐먼트 안에 객체의 형태로 도큐먼트가 있는 경우

	// queryString -> default하게 query값을 줄 떄

	//sortPipeLine := bson.D{
	// {"$sort", bson.D{
	//    {"reg", 1},
	// }},
	//}
	//
	//lookupStage := []bson.D{
	// {
	//    {"$lookup", bson.D{
	//       {"from", m.nftCollectibles.Name()}, // 대상 컬렉션 이름으로 변경해야 함
	//       {"let", bson.D{
	//          {"address", "$address"},
	//       }},
	//       {"pipeline", bson.A{
	//          bson.D{
	//             {"$match", bson.D{
	//                {"$expr", bson.D{
	//                   {"$and", bson.A{
	//                      bson.D{
	//                         {"$eq", bson.A{
	//                            "$address",
	//                            "$$address",
	//                         }},
	//                      },
	//                   }},
	//                }},
	//             }},
	//          },
	//       }},
	//       {"as", "joinedData"}, // 조인 결과를 저장할 필드 이름을 지정해야 함
	//    }},
	// },
	// {
	//    {"$unwind", "$joinedData"},
	//"preserveNullAndEmptyArrays": true,
	// },
	//}

	// Aggregate pipeline
	pipeline := mongo.Pipeline{
		searchStage,
		//sortPipeLine,
	}

	//pipeline = append(pipeline, lookupStage...)
	//
	//pagePipeLine := []bson.D{
	// {
	//    {"$skip", 10},
	// },
	// {
	//    {"$limit", 5},
	// },
	//}
	//
	//pipeline = append(pipeline, pagePipeLine...)

	cursor, err := m.search.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(results))

	// Process the results
	for _, result := range results {
		fmt.Println(result)
	}

}
