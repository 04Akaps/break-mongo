package network

import (
	"break-mongo/document"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Network struct {
	engine        *gin.Engine
	document      *document.Document
	shardDocument *document.ShardingDocument
}

func NewNetwork() *Network {
	r := &Network{
		engine:        gin.New(),
		document:      document.NewDocument(),
		shardDocument: document.NewShardingDocument(),
	}

	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	NewUserAPI(r, r.document)
	NewBulkUserAPI(r, r.document)
	NewFileAPI(r, r.document)
	NewAggregateAPI(r, r.document)
	NewShardAPI(r, r.shardDocument, r.document)
	NewSearchAPI(r, r.document)

	return r
}

func (r *Network) Run(port string) error {
	return r.engine.Run(port)
}
