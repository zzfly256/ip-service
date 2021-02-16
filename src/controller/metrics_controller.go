package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/ip-service/src/helper"
	"net/http"
	"strconv"
)

func GetServiceMetrics(context *gin.Context) {
	ipData := helper.LoadIpData()

	// IP 库总条目数指标
	str := "ip_service_data_count{service=\"ip-service\"} " + strconv.Itoa(len(ipData.List))
	// 当前进程内缓存 IP 数指标
	str += "\nip_service_cache_count{service=\"ip-service\"} " + strconv.Itoa(len(ipData.Cache))

	context.String(http.StatusOK, str)
}
