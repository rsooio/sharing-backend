package response

import (
	"net/http"
)

const defaultCode = http.StatusBadRequest

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func FromCodeError(code int, err error) error {
	if err == nil {
		return nil
	}
	return NewCodeError(code, err.Error())
}

func FromDefaultError(err error) error {
	if err == nil {
		return nil
	}
	return NewDefaultError(err.Error())
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
