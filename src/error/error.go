package error

import (
	"fmt"
)

type StandardError struct {
	Code    int
	Message string
}

var (
	IpNotFound = StandardError{30001, "IP地址获取失败"}
)

func (error StandardError) Error() string {
	return fmt.Sprintf("[%d] %s", error.Code, error.Message)
}
