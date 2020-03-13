package errors

import (
	"github.com/micro/go-micro/errors"
	"my-micro/demo/src/share/config"
)

const (
	errorCodeCommentSuccess = 200
)

var (
	ErrorCommentFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeCommentSuccess,
	)
)

