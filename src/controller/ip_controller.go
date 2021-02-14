package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/IpService/src/error"
	"github.com/zzfly256/IpService/src/helper"
)

func QueryMyIp(context *gin.Context) {
	ip := context.ClientIP()
	if len(ip) == 0 {
		context.Set("error", error.IpNotFound)
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

	// 匹配结果
	data := helper.GetIpInfo(ip)

	// 产生响应
	context.Set("data", data)
}
