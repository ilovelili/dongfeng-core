package utils

import (
	"github.com/micro/go-micro/errors"
)

// NewError new rpc error
func NewError(detail string, httpcode ...int32) error {
	var code int32
	if len(httpcode) == 0 {
		code = errors.Parse(detail).Code
	} else {
		code = httpcode[0]
	}

	return errors.New("dongfeng.svc.core.server", detail, code)
}
