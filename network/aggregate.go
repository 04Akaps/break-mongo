package network

import (
	"break-mongo/document"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	aggregateApiInit     sync.Once
	aggregateApiInstance *aggregateAPI
)

type aggregateAPI struct {
	router *Network
	doc    *document.Document
}

func NewAggregateAPI(router *Network, doc *document.Document) *aggregateAPI {
	aggregateApiInit.Do(func() {
		aggregateApiInstance = &aggregateAPI{
			router: router,
			doc:    doc,
		}

		router.RegisterPOSTHandler(BASE_AGGREAGTE_URI+"/dummy", aggregateApiInstance.dummy)

		router.RegisterGETHandler(BASE_AGGREAGTE_URI+"/findAllOne", aggregateApiInstance.findAllOne)
		router.RegisterGETHandler(BASE_AGGREAGTE_URI+"/findAllTwo", aggregateApiInstance.findAllTwo)
		router.RegisterGETHandler(BASE_AGGREAGTE_URI+"/byName/:name", aggregateApiInstance.findByName)

	})

	return aggregateApiInstance
}

func (a *aggregateAPI) dummy(c *gin.Context) {
	a.doc.AggregateDummy()
}

func (a *aggregateAPI) findAllOne(c *gin.Context) {
	result, _ := a.doc.FindAllAggregateOne()
	a.router.RespOK(c, result)
}

func (a *aggregateAPI) findAllTwo(c *gin.Context) {
	result, _ := a.doc.FindAllAggregateTwo()
	a.router.RespOK(c, result)
}

func (a *aggregateAPI) findByName(c *gin.Context) {
	name := c.Param("name")

	result, _ := a.doc.FindByNameAggregate(name)

	a.router.RespOK(c, result)

}
