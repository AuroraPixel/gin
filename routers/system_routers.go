package routers

import (
	"gin/api"

	"github.com/gin-gonic/gin"
)

func (r Routers) SystemRouter(c *gin.RouterGroup) {
	systemapi := api.GroupApi.SystemApi
	c.GET("", systemapi.SystemApiVO)
}
