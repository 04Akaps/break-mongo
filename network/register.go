package network

import "github.com/gin-gonic/gin"

func (r *Network) RegisterPOSTHandler(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.POST(relativePath, handlers...)
}

func (r *Network) RegisterGETHandler(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.GET(relativePath, handlers...)
}

func (r *Network) RegisterPUTHandler(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.PUT(relativePath, handlers...)
}
