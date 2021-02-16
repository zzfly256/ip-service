package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/ip-service/src/error"
	"github.com/zzfly256/ip-service/src/helper"
)

func QueryMyIp(context *gin.Context) {
	ip := context.ClientIP()
	if len(ip) == 0 {
		context.Set("error", error.IpParameterError)
	} else {
		context.Set("data", ip)
	}
}

func QueryIpAddress(context *gin.Context) {
	// 获取请求的 IP
	ip := context.Query("ip")
	if len(ip) == 0 {
		ip = context.PostForm("ip")
	}
	if len(ip) == 0 {
		ip = context.ClientIP()
	}
	if !helper.CheckIp(ip) {
		context.Set("error", error.IpParameterError)
		return
	}

	// 匹配结果
	data := helper.GetIpInfo(ip)

	// 产生响应
	context.Set("data", data)
}
