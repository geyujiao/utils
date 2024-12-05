package api

import (
	"fmt"
	"net/http"
	"utils/request"

	"github.com/gin-gonic/gin"
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
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	fmt.Println("PostJsonJ param", param)
	c.JSON(http.StatusOK, "success")
	return
}

// post form-data
func PostFormData(c *gin.Context) {
	mobile, b := c.GetPostForm("mobile")
	if !b {
		fmt.Println("GetPostForm mobile b", b)
		c.JSON(http.StatusOK, "fail")
		return
	}

	fmt.Println("GetPostForm mobile", mobile)
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

	fmt.Println(count)
	fmt.Println("mobileGet", mobileGet)
	fmt.Println("mobileGetString", mobileGetString)
	fmt.Println("mobileGetQueryArray", mobileGetQueryArray)
	fmt.Println("mobileGetQueryMap", mobileGetQueryMap)
	fmt.Println("mobileQuery", mobileQuery)
	fmt.Println("mobileParam", mobileParam)
*/

// Get
func Get(c *gin.Context) {
	mobileGetQuery, _ := c.GetQuery("mobile")

	fmt.Println("mobileGetQuery", mobileGetQuery)

	c.JSON(http.StatusOK, "success")
	return
}

// Get
func GetServerStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "get /server-status success")
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
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, "PostFormFile err fail")
		return
	}

	c.JSON(http.StatusOK, "PostFormFile---"+string(resp))

	return
}
