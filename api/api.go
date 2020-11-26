package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vgmdj/utils/logger"
	"net/http"
	"utils/request"
)

type PostJsonReq struct {
	Mobile string `json:"Mobile"`
	Count  int    `json:"Count"`
	Tel    string `json:"Tel"`
}

// Post
func PostJson(c *gin.Context) {
	param := PostJsonReq{}
	if err := c.BindJSON(&param); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	logger.Info("PostJsonJ param", param)
	c.JSON(http.StatusOK, "success")
	return
}

// post form-data
func PostFormData(c *gin.Context) {
	mobile, b := c.GetPostForm("mobile")
	if !b {
		logger.Error("GetPostForm mobile b", b)
		c.JSON(http.StatusOK, "fail")
		return
	}

	logger.Info("GetPostForm mobile", mobile)
	c.JSON(http.StatusOK, "success")
	return
}

/*
	count := c.GetInt("count")               // 不能用
	mobileGet, _ := c.Get("mobile")          // 不能用
	mobileGetString := c.GetString("mobile") // 不能用
	mobileGetQueryArray, _ := c.GetQueryArray("mobile")
	mobileGetQueryMap, _ := c.GetQueryMap("mobile") // 不能用
	mobileQuery := c.Query("mobile")                // === GetQuery
	mobileParam := c.Param("mobile")                // 不能用

	logger.Info(count)
	logger.Info("mobileGet", mobileGet)
	logger.Info("mobileGetString", mobileGetString)
	logger.Info("mobileGetQueryArray", mobileGetQueryArray)
	logger.Info("mobileGetQueryMap", mobileGetQueryMap)
	logger.Info("mobileQuery", mobileQuery)
	logger.Info("mobileParam", mobileParam)
*/

// Get
func Get(c *gin.Context) {
	mobileGetQuery, _ := c.GetQuery("mobile")

	logger.Info("mobileGetQuery", mobileGetQuery)

	c.JSON(http.StatusOK, "success")
	return
}

// 测试request.PostFormDataFile
func PostFormDataFile(c *gin.Context) {
	file, _ := c.FormFile("video")

	if file == nil {
		c.JSON(http.StatusOK, "PostFormDataFile file is nil")
		return
	}

	c.JSON(http.StatusOK, "PostFormDataFile success")

	return
}

// postman 传 file
func PostFormFile(c *gin.Context) {
	file, _ := c.FormFile("video")

	if file == nil {
		c.JSON(http.StatusOK, "PostFormFile file is nil")
		return
	}

	resp, err := request.PostFormDataFile("http://127.0.0.1:83/postformdatafile", file)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, "PostFormFile err fail")
		return
	}

	c.JSON(http.StatusOK, "PostFormFile---" + string(resp))

	return
}
