package defs

import (
	"fmt"
)

type BasicError struct {
	ErrorMsg   error
	ErrorCode  string
}

type Error struct {
	ErrorDetail     BasicError
	HttpCode        int
}

var (
	RequestParamError = Error {
		ErrorDetail: BasicError {
			ErrorMsg: fmt.Errorf("request param have mistakes"),
			ErrorCode: "001",
		},
		HttpCode: 400,
	}
	AuthenticateError = Error {
		ErrorDetail: BasicError {
			ErrorMsg: fmt.Errorf("authenticate failed"),
			ErrorCode: "002",
		},
		HttpCode: 401,
	}
)