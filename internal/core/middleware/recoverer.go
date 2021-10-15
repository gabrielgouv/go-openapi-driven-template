package middleware

import (
	"encoding/json"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/core/err"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicErr := recover()
			if panicErr != nil {
				// TODO: Log panic error
				bodyErr := err.NewApiErrorBuilder().
					Code("INTERNAL_SERVER_ERROR").
					Message("Internal Server Error").
					MessageDetail("An unrecoverable error has occurred").
					Build()
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(bodyErr)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
