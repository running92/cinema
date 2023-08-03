package errors

import (
	"cinema/src/share/config"
	"go-micro.dev/v4/errors"
)

const (
	errorCodeOrderSuccess = 200
)

var (
	ErrorOrderFailed = errors.New(
		config.ServiceNameOrder, "操作异常", errorCodeOrderSuccess,
	)
	ErrorOrderAlreadyScore = errors.New(
		config.ServiceNameOrder, "已经评分了！", errorCodeOrderSuccess,
	)
)
