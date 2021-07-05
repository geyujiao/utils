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
//
//func main()  {
//	reader := &study.ReadFile{
//		Path: "./study/log.txt",
//	}
//	writer := &study.WriteInfluxDB{
//		InfluxDB: "",
//	}
//
//	log := &study.LogProcess{
//		Rc: make(chan []byte, 1),
//		Wc: make(chan string, 1),
//		Read: reader,
//		Write: writer,
//	}
//	go log.Read.Read(log.Rc)
//	go log.Process()
//	go log.Write.Write(log.Wc)
//
//	time.Sleep(30 *time.Second)
//}
