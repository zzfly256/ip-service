package helper

import (
	"bufio"
	"fmt"
	"github.com/zzfly256/ip-service/src/idl"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type IpData struct {
	List  idl.IpAddressInfoList
	Cache map[string]idl.IpAddressInfoItem
}

var ipData *IpData
var cacheLock sync.RWMutex

// 加载 IP 地址库
func LoadIpData() *IpData {
	// TODO: 二重锁 or 饿汉模式
	if ipData == nil {
		path := filepath.Dir(os.Args[0]) + "/storage/czip.txt"
		fp, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		defer fp.Close()

		ipData = &IpData{
			List: []idl.IpAddressInfoItem{},
		}

		reader := bufio.NewReader(fp)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				fmt.Println("IP 地址加载完成")
				break
			}
			if err != nil {
				panic(err)
			}

			lineSlice := strings.Fields(line)

			// 非预期格式则跳过
			if len(lineSlice) < 3 {
				continue
			}

			ipItem := idl.IpAddressInfoItem{
				Area:         lineSlice[2],
				Isp:          strings.Trim(fmt.Sprint(lineSlice[3:]), "[]"),
				SegmentStart: transformIpAddressToStdString(lineSlice[0]),
				SegmentEnd:   transformIpAddressToStdString(lineSlice[1]),
			}
			ipData.List = append(ipData.List, ipItem)
		}

		ipData.Cache = make(map[string]idl.IpAddressInfoItem)
	}
	return ipData
}

// 匹配 IP 地址信息
func GetIpInfo(ipAddress string) idl.IpAddressInfoItem {

	ipData := LoadIpData()

	cacheLock.RLock()

	if ipData.Cache[ipAddress].Ip == "" {
		cacheLock.RUnlock()

		ipTarget := transformIpAddressToStdString(ipAddress)
		targetItem := idl.IpAddressInfoItem{}

		for _, item := range ipData.List {
			if strings.Compare(ipTarget, item.SegmentStart) >= 0 {
				if strings.Compare(ipTarget, item.SegmentEnd) <= 0 {
					item.Ip = ipAddress
					targetItem = item
					break
				}
			}
		}

		cacheLock.Lock()
		ipData.Cache[ipAddress] = targetItem
		cacheLock.Unlock()
		cacheLock.RLock()
	}

	result := ipData.Cache[ipAddress]
	cacheLock.RUnlock()

	return result
}

// 将 IP 地址转化为统一长度的字符串
func transformIpAddressToStdString(ipAddress string) string {
	ipSlice := strings.Split(ipAddress, ".")
	result := ""

	for _, ipSegmentRawStr := range ipSlice {
		ipSegmentInteger, err := strconv.Atoi(strings.Trim(ipSegmentRawStr, "\ufeff"))
		if err != nil {
			panic(err)
		}

		ipSegmentStr := fmt.Sprintf("%03d", ipSegmentInteger)
		result = result + ipSegmentStr
	}

	return result
}

func CheckIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}
