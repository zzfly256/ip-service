package error

import (
	"fmt"
)

type StandardError struct {
	Code    int
	Message string
}

var (
	IpParameterError = StandardError{30001, "IP地址参数错误"}
)

func (error StandardError) Error() string {
	return fmt.Sprintf("[%d] %s", error.Code, error.Message)
}
