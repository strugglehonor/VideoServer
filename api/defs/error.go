package defs

import (
	"fmt"
	"net/http"
)

type BasicError struct {
	ErrorMsg  error
	ErrorCode string
}

type Error struct {
	ErrorDetail BasicError
	HttpCode    int
}

var (
	RequestParamError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("request param have mistakes"),
			ErrorCode: "001",
		},
		HttpCode: http.StatusBadRequest,
	}
	AuthenticateError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("authenticate failed"),
			ErrorCode: "002",
		},
		HttpCode: http.StatusUnauthorized,
	}
	UnmarshalError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("encoding/json unmarshal failed"),
			ErrorCode: "003",
		},
		HttpCode: http.StatusInternalServerError,
	}
	MarshalError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("encoding/json marshal failed"),
			ErrorCode: "004",
		},
		HttpCode: http.StatusInternalServerError,
	}
	DBInsertError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("Insert data to DB Failed"),
			ErrorCode: "005",
		},
		HttpCode: http.StatusInternalServerError,
	}
	DBSelectError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("Select data from DB Failed"),
			ErrorCode: "005",
		},
		HttpCode: http.StatusInternalServerError,
	}
	ErrorInternalFaults = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("Internal Failed"),
			ErrorCode: "006",
		},
		HttpCode: http.StatusInternalServerError,
	}
	SessionExpiredError = Error{
		ErrorDetail: BasicError{
			ErrorMsg:  fmt.Errorf("session expired"),
			ErrorCode: "007",
		},
		HttpCode: http.StatusForbidden,
	}
)
