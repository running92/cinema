package errors

import (
	"cinema/src/share/config"
	"go-micro.dev/v4/errors"
)

const (
	errorCodeFilmSuccess = 200
)

var (
	ErrorFilmFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeFilmSuccess,
	)
)
