package err

import (
	"fmt"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/core/util/dateutil"
)

type ApiError struct {
	HttpStatusCode int    `json:"-"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
	MessageDetail  string `json:"messageDetail,omitempty"`
	At             string `json:"at,omitempty"`
	Err            error  `json:"-"`
}

type ApiErrorBuilder interface {
	HttpStatusCode(code int) ApiErrorBuilder
	Code(code string) ApiErrorBuilder
	Message(message string) ApiErrorBuilder
	MessageDetail(messageDetail string) ApiErrorBuilder
	Err(error error) ApiErrorBuilder
	Build() ApiError
}

type apiErrorBuilder struct {
	httpStatusCode int
	code           string
	message        string
	messageDetail  string
	err            error
}

func (aeb *apiErrorBuilder) HttpStatusCode(code int) ApiErrorBuilder {
	aeb.httpStatusCode = code
	return aeb
}

func (aeb *apiErrorBuilder) Code(code string) ApiErrorBuilder {
	aeb.code = code
	return aeb
}

func (aeb *apiErrorBuilder) Message(message string) ApiErrorBuilder {
	aeb.message = message
	return aeb
}

func (aeb *apiErrorBuilder) MessageDetail(messageDetail string) ApiErrorBuilder {
	aeb.messageDetail = messageDetail
	return aeb
}

func (aeb *apiErrorBuilder) Err(error error) ApiErrorBuilder {
	aeb.err = error
	return aeb
}

func (aeb *apiErrorBuilder) Build() ApiError {
	if aeb.httpStatusCode < 100 {
		aeb.httpStatusCode = 500
	}
	if len(aeb.message) <= 0 {
		aeb.message = "An error occurred"
	}
	return ApiError{
		HttpStatusCode: aeb.httpStatusCode,
		Code:           aeb.code,
		Message:        aeb.message,
		MessageDetail:  aeb.messageDetail,
		Err:            aeb.err,
		At:             dateutil.NowAsString(),
	}
}

func (r *ApiError) Error() string {
	return fmt.Sprintf("%s: message: %s, detail: %s", r.Code, r.Message, r.MessageDetail)
}

func NewApiErrorBuilder() ApiErrorBuilder {
	return &apiErrorBuilder{}
}
