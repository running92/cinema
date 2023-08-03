package errors

import (
	"cinema/src/share/config"
	"go-micro.dev/v4/errors"
)

const (
	errorCodeCinemaSuccess = 200
)

var (
	ErrorCinemaFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeCinemaSuccess,
	)
	ErrorCinemaNotFound = errors.New(
		config.ServiceNameUser, "找不到对应的影院", errorCodeCinemaSuccess,
	)
)
