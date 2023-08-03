package errors

import (
	"cinema/src/share/config"
	"go-micro.dev/v4/errors"
)

const (
	errorCodeCommentSuccess = 200
)

var (
	ErrorCommentFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeCommentSuccess,
	)
)
