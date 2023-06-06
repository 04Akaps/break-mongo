package network

import (
	"break-mongo/document"
	"break-mongo/types"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

var (
	userApiInit     sync.Once
	userApiInstance *userAPI
)

type userAPI struct {
	router *Network
	doc    *document.Document
}

func NewUserAPI(router *Network, doc *document.Document) *userAPI {
	userApiInit.Do(func() {
		userApiInstance = &userAPI{
			router: router,
			doc:    doc,
		}

		userApiInstance.router.RegisterPOSTHandler(BASE_USER_URI+"/dummy", userApiInstance.registerDummyUserData)
		userApiInstance.router.RegisterPUTHandler(BASE_USER_URI+"/update", userApiInstance.updateUserData)
		userApiInstance.router.RegisterGETHandler(BASE_USER_URI+"/findAll", userApiInstance.findUserData)
		userApiInstance.router.RegisterGETHandler(BASE_USER_URI+"/find", userApiInstance.findOneUserData)
		userApiInstance.router.RegisterGETHandler(BASE_USER_URI+"/sortFind", userApiInstance.findUserDataByAgeSort)

	})

	return userApiInstance
}

func (u *userAPI) registerDummyUserData(c *gin.Context) {
	log.Println("Test용 더미 유저 데이터 등록")
	log.Println("원래는 Insert Many같은 함수를 사용, 어차피 더미니깐 그냥 내맘대로 작성")

	for i := 1; i < 50; i++ {

		dummyUser := &types.User{
			Age: int64(i + 1),
		}

		if i < 10 {
			dummyUser.Name = "A"
		} else if i < 20 {
			dummyUser.Name = "B"
		} else if i < 30 {
			dummyUser.Name = "C"
		} else {
			dummyUser.Name = "D"
		}

		u.doc.InsertDummyUserData(dummyUser)
	}
}

func (u *userAPI) updateUserData(c *gin.Context) {
	var req types.User

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.Resp(c, 201, "Bind Error")
		return
	}

	if err := u.doc.UpdateUserData(&req); err != nil {
		log.Println(err)
		u.router.Resp(c, 401, "Update Failed")
	} else {
		log.Println(err)
		u.router.RespOK(c, "Success UPdated")
	}

}

func (u *userAPI) findUserData(c *gin.Context) {
	// Find ALl
	if data, err := u.doc.FindUserData(); err != nil {
		// 상황에 따른 에러처리는 알아서;;
		log.Println(err)
		u.router.Resp(c, 201, "Not Found")
	} else {
		u.router.RespOK(c, data)
	}
}

func (u *userAPI) findOneUserData(c *gin.Context) {
	// FindOne
	var req types.User

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		u.router.Resp(c, 201, "Bind Error")
		return
	}

	if data, err := u.doc.FindOneUserData(&req); err != nil {
		// 상황에 따른 에러처리는 알아서;;
		log.Println(err)
		u.router.Resp(c, 201, "Not Found")
	} else {
		u.router.RespOK(c, data)
	}
}

func (u *userAPI) findUserDataByAgeSort(c *gin.Context) {
	var req types.User

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.Resp(c, 201, "Bind Error")
		return
	}

	if data, err := u.doc.FindUserDataByAgeSort(); err != nil {
		log.Println(err)
		u.router.Resp(c, 201, "Not Found")
	} else {
		u.router.RespOK(c, data)
	}
}
