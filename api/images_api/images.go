package imagesapi

import imageservice "gin/service/image_service"

type ImagesApi struct {
	ImageService imageservice.ImageService
}

// 符合场景的图片格式
var CheckImageType = []string{
	"jpg",
	"png",
	"tif",
	"gif",
	"bmp",
	"svg",
}

// 图片返回视图
type ImagesVO struct {
	Path      string `json:"path"`      //路径
	FileName  string `json:"fileName"`  //文件名
	IsSuccess bool   `json:"isSuccess"` //是否成功
	Message   string `json:"message"`   //消息内容
}
