package middleware

import (
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/pkg/properties"
	"github.com/go-chi/httplog"
	"net/http"
)

func HttpLogger(next http.Handler) http.Handler {
	serviceName := properties.GetInstance().Name
	httpLogger := httplog.NewLogger(serviceName, httplog.Options{
		JSON: false,
	})
	return httplog.RequestLogger(httpLogger)(next)
}
