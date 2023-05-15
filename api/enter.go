package api

import (
	echoapi "gin/api/echo_api"
	imagesapi "gin/api/images_api"
	systemapi "gin/api/system_api"
)

type ApiGroup struct {
	SystemApi systemapi.SystemApi
	ImagesApi imagesapi.ImagesApi
	EchoApi   echoapi.EchoApi
}

var GroupApi = new(ApiGroup)
