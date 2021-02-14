package idl

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IpAddressInfoItem struct {
	Ip           string `json:"ip"`
	Area         string `json:"area"`
	Isp          string `json:"isp"`
	SegmentStart string `json:"-"`
	SegmentEnd   string `json:"-"`
}

type IpAddressInfoList []IpAddressInfoItem
