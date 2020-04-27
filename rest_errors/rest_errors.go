package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status()  int
	Error()   string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"ErrMessage"`
	ErrStatus  int           `json:"ErrStatus"`
	ErrError   string        `json:"ErrError"`
	ErrCauses  []interface{} `json:"ErrCauses"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("ErrMessage: %s - ErrStatus: %d - ErrError: %s - ErrCauses: %v ",
			e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string{
	return e.ErrMessage
}

func (e restErr) Status()  int{
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, status int, error string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   error,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error){
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr) ; err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr{
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server ErrError",
	}
	if err!=nil{
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}


