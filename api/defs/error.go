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
	UnmarshalError = Error {
		ErrorDetail: BasicError{
			ErrorMsg: fmt.Errorf("encoding/json unmarshal failed"),
			ErrorCode: "003",
		},
		HttpCode: 500,
	}
	MarshalError = Error {
		ErrorDetail: BasicError{
			ErrorMsg: fmt.Errorf("encoding/json marshal failed"),
			ErrorCode: "004",
		},
		HttpCode: 500,
	}
	DBInsertError = Error {
		ErrorDetail: BasicError{
			ErrorMsg: fmt.Errorf("Insert data to DB Failed"),
			ErrorCode: "005",
		},
		HttpCode: 500,
	}
	ErrorInternalFaults = Error {
		ErrorDetail: BasicError{
			ErrorMsg: fmt.Errorf("Internal Failed"),
			ErrorCode: "006",
		},
		HttpCode: 500,
	}
	SessionExpiredError = Error {
		ErrorDetail: BasicError{
			ErrorMsg: fmt.Errorf("session expired"),
			ErrorCode: "007",
		},
		HttpCode: 403,
	}
)