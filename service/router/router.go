package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

	}

	return r
}
