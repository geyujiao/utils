package router

import (
	"github.com/gin-gonic/gin"
	"utils/api"
)

func Register(r *gin.Engine) {

	/*
		request
	*/
	r.POST("/postjson", api.PostJson)                 // post json
	r.POST("/formdata", api.PostFormData)             // post formdata
	r.POST("/postformdatafile", api.PostFormDataFile) // post formdata(file)
	r.GET("/get", api.Get)                            // get
	r.POST("/postformdfile", api.PostFormFile)        // postman 传入file,再调用接口 /postformdatafile

}
