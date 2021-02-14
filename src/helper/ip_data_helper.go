package helper

import (
	"bufio"
	"fmt"
	"github.com/zzfly256/IpService/src/idl"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type IpData struct {
	list idl.IpAddressInfoList
}

var ipData *IpData

// 加载 IP 地址库
func LoadIpData() *IpData {
	// TODO: 二重锁 or 饿汉模式
	if ipData == nil {
		path := filepath.Dir(os.Args[0]) + "/storage/czip.txt"
		// 测试
		path = "/mnt/c/Users/rytia/Desktop/LearningGo/storage/czip.txt"
		fp, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		defer fp.Close()

		ipData = &IpData{
			list: []idl.IpAddressInfoItem{},
		}

		reader := bufio.NewReader(fp)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			lineSlice := strings.Fields(line)

			ipItem := idl.IpAddressInfoItem{
				Area:         lineSlice[2],
				Isp:          strings.Trim(fmt.Sprint(lineSlice[3:]), "[]"),
				SegmentStart: transformIpAddressToStdString(lineSlice[0]),
				SegmentEnd:   transformIpAddressToStdString(lineSlice[1]),
			}
			ipData.list = append(ipData.list, ipItem)
		}
	}
	return ipData
}

// 匹配 IP 地址信息
func GetIpInfo(ipAddress string) idl.IpAddressInfoItem {

	ipData := LoadIpData()
	ipTarget := transformIpAddressToStdString(ipAddress)

	for _, item := range ipData.list {
		if strings.Compare(ipTarget, item.SegmentStart) >= 0 {
			if strings.Compare(ipTarget, item.SegmentEnd) <= 0 {
				item.Ip = ipAddress
				return item
			}
		}
	}

	result := idl.IpAddressInfoItem{}

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
