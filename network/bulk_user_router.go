package network

import (
	"break-mongo/document"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

var (
	bulkUserApiInit     sync.Once
	bulkUserApiInstance *bulkUserAPI
)

type bulkUserAPI struct {
	router *Network
	doc    *document.Document
}

func NewBulkUserAPI(router *Network, doc *document.Document) *bulkUserAPI {
	bulkUserApiInit.Do(func() {
		bulkUserApiInstance = &bulkUserAPI{
			router: router,
			doc:    doc,
		}

		router.RegisterPOSTHandler(BASE_BULK_USER_URI+"/dummy", bulkUserApiInstance.registerBulkUserData)

		router.RegisterGETHandler(BASE_BULK_USER_URI+"/findAll", bulkUserApiInstance.findAllBulkUserData)
		router.RegisterGETHandler(BASE_BULK_USER_URI+"/findAllBySort", bulkUserApiInstance.findAllBulkUserBySort)

		router.RegisterGETHandler(BASE_BULK_USER_URI+"/addIndex", bulkUserApiInstance.addNameIndex)
		router.RegisterGETHandler(BASE_BULK_USER_URI+"/addDoubleIndex", bulkUserApiInstance.addDoubleIndex)

	})

	return bulkUserApiInstance
}

func (u *bulkUserAPI) registerBulkUserData(c *gin.Context) {
	// 많은 데이터를 한번에 삽입 할 떄,
	u.doc.InsertBulkDummyUserData()
}

func (u *bulkUserAPI) findAllBulkUserData(c *gin.Context) {

	startTime := time.Now()

	if result, err := u.doc.FindAllBulkUser(true); err != nil {

		log.Println(err)
		u.router.Resp(c, 401, "Find ALl Failed")
	} else {

		elapsedTime := time.Since(startTime) // 총 걸린 시간 계산
		log.Println("쿼리 실행 시간: ", elapsedTime)

		u.router.RespOK(c, result)
	}
}

func (u *bulkUserAPI) findAllBulkUserBySort(c *gin.Context) {

	startTime := time.Now()

	if result, err := u.doc.FindAllBulkUserBySort(); err != nil {

		log.Println(err)
		u.router.Resp(c, 401, "Find ALl Failed")
	} else {

		elapsedTime := time.Since(startTime) // 총 걸린 시간 계산
		log.Println("쿼리 실행 시간: ", elapsedTime)

		u.router.RespOK(c, result)
	}
}

func (u *bulkUserAPI) addNameIndex(c *gin.Context) {
	u.doc.AddIndexByName()
}

func (u *bulkUserAPI) addDoubleIndex(c *gin.Context) {
	u.doc.AddDoubleIndex()
}
