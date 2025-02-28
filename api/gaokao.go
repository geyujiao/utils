package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsbGetAdmissionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
	return
}

func JsbGetScoreInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
	return
}

func GetVerifyCode(c *gin.Context) {
	c.Header("resp-code", "4000")
	c.JSON(http.StatusOK, "success")
	fullPath := c.FullPath() // 获取完整的路径，例如 /testpath?query=value
	println("------" + fullPath)
	return
}

func YunchenOptions(c *gin.Context) {
	c.JSON(http.StatusOK, "options success")
	return
}
