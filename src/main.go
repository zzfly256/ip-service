package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/IpService/src/controller"
	"github.com/zzfly256/IpService/src/helper"
	"github.com/zzfly256/IpService/src/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	helper.LoadIpData()

	v1 := router.Group("/v1").Use(middleware.FormatResponse())
	{
		v1.GET("/query_my_ip", controller.QueryMyIp)
		v1.GET("/query_ip_address", controller.QueryIpAddress)
		v1.POST("/query_ip_address", controller.QueryIpAddress)
		v1.POST("/reload_ip_data", controller.QueryIpAddress)
	}

	router.GET("/get_metrics", controller.GetServiceMetrics)

	_ = router.Run(":3000")
}
