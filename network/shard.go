package network

import (
	"break-mongo/document"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	shardApiInit     sync.Once
	shardApiInstance *shardAPI
)

type shardAPI struct {
	router        *Network
	shardDocument *document.ShardingDocument
	d             *document.Document
}

func NewShardAPI(router *Network, doc *document.ShardingDocument, d *document.Document) *shardAPI {
	shardApiInit.Do(func() {
		shardApiInstance = &shardAPI{
			router:        router,
			shardDocument: doc,
			d:             d,
		}

	})

	router.RegisterPOSTHandler(BASE_SHARD_URI+"/addDummy", shardApiInstance.addShardDummy)

	return shardApiInstance
}

func (s *shardAPI) addShardDummy(c *gin.Context) {
	s.shardDocument.AddShardDummy()
}
