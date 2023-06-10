package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Network) Resp(c *gin.Context, status int, resp interface{}) {
	c.JSON(status, resp)
}

func (r *Network) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

const (
	BASE_USER_URI      = "/user"
	BASE_BULK_USER_URI = "/bulk-user"
	BASE_FILE_URI      = "/file"
	BASE_AGGREAGTE_URI = "/aggregate"
)
