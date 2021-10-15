package handler

import (
	"github.com/gabrielgouv/go-openapi-driven-template/generated/openapi"
	"github.com/gabrielgouv/go-openapi-driven-template/internal/core/err"
	"github.com/gabrielgouv/go-openapi-driven-template/internal/core/factory"
	"net/http"
)

func (p *Handler) Echo(w http.ResponseWriter, _ *http.Request, params openapi.EchoParams) {
	if params.Message == "error" {
		myError := err.NewApiErrorBuilder().
			HttpStatusCode(400).
			Code("ERR00001").
			Message("An error occurred").
			MessageDetail("Cannot get echo").
			Build()
		p.Error(w, myError)
		return
	}
	if params.Message == "panic" {
		panic("Panic!")
	}
	echoService := factory.MakeEchoService()
	echo := echoService.GetEcho(params.Message)
	p.Json(w, openapi.Echo{
		Message:      &echo.Message,
		Date:         &echo.Date,
		RandomNumber: &echo.RandomNumber,
	})
}
