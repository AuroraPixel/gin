package routers

import (
	"gin/api"

	"github.com/gin-gonic/gin"
)

func (r Routers) imagesRouter(c *gin.RouterGroup) {
	ImagesApi := api.GroupApi.ImagesApi
	ImagesApi.InitController()
	//图片文件上传
	c.POST("images/upload", ImagesApi.Upload)
	//分页查询列表
	c.GET("images/page", ImagesApi.PageList)
}
