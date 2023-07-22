package network

import (
	"break-mongo/document"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	searchApiInit sync.Once
	searchApi     *searchAPI
)

type searchAPI struct {
	router *Network
	d      *document.Document
}

func NewSearchAPI(router *Network, d *document.Document) *searchAPI {
	searchApiInit.Do(func() {
		searchApi = &searchAPI{
			router: router,
			d:      d,
		}

		router.RegisterGETHandler("/search", searchApi.search)
	})

	return searchApi
}

func (s *searchAPI) search(c *gin.Context) {
	s.d.SearchTest()
}
