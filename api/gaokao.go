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
	return
}

func YunchenOptions(c *gin.Context) {
	c.JSON(http.StatusOK, "options success")
	return
}
