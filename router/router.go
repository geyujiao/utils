package router

import (
	"utils/api"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	/*
		request
	*/
	r.POST("/postjson", api.PostJson)                 // post json
	r.POST("/formdata", api.PostFormData)             // post formdata
	r.POST("/postformdatafile", api.PostFormDataFile) // post formdata(file)
	r.GET("/get", api.Get)                            // get
	r.GET("/server-status", api.GetServerStatus)      // get
	r.POST("/postformdfile", api.PostFormFile)        // postman 传入file,再调用接口 /postformdatafile
	r.POST("/api/v1/jsbGetAdmissionInfo", api.JsbGetAdmissionInfo)
	r.POST("/api/v1/jsbGetScoreInfo", api.JsbGetScoreInfo)
	r.POST("/api/v1/getVerifyCode", api.GetVerifyCode)
	r.OPTIONS("/yunchen", api.YunchenOptions)
	r.POST("/openapi/httpService/ISSSService", api.GetVerifyCode)

}
