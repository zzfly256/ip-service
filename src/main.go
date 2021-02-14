package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/IpService/src/controller"
	"github.com/zzfly256/IpService/src/middleware"
)

func main() {
	router := gin.Default()
	router.Use(middleware.FormatResponse())

	v1 := router.Group("/v1")
	{
		v1.GET("/query_my_ip", controller.QueryMyIp)
		v1.GET("/query_ip_address", controller.QueryIpAddress)
		v1.POST("/query_ip_address", controller.QueryIpAddress)
		v1.POST("/reload_ip_data", controller.QueryIpAddress)
	}

	_ = router.Run(":3000")
}
