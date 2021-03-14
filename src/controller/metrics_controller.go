package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/ip-service/src/helper"
	"net/http"
	"runtime"
	"strconv"
)

func GetServiceMetrics(context *gin.Context) {
	ipData := helper.LoadIpData()

	// IP 库总条目数指标
	str := "ip_service_data_num{service=\"ip-service\"} " + strconv.Itoa(len(ipData.List))
	// 当前进程内缓存 IP 数指标
	str += "\nip_service_cache_num{service=\"ip-service\"} " + strconv.Itoa(len(ipData.Cache))
	// 当前程序的协程数
	str += "\nip_service_cpu_num{service=\"ip-service\"} " + strconv.Itoa(runtime.NumCPU())
	str += "\nip_service_goroutine_num{service=\"ip-service\"} " + strconv.Itoa(runtime.NumGoroutine())

	context.String(http.StatusOK, str)
}
