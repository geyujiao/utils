package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vgmdj/utils/logger"
	"time"
	"utils/router"
)

func main()  {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Register(r)

	address := fmt.Sprintf(":83")
	err := r.Run(address)
	if err != nil {
		logger.Error("start http server error:", err.Error())
		return
	}
}
