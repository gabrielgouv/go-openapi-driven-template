package handler

import (
	"encoding/json"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/generated/openapi"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/core/err"
	"net/http"
)

type Handler struct{}

var _ openapi.ServerInterface = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

func (p *Handler) Json(w http.ResponseWriter, v interface{}) {
	p.AddHeader(w, "Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(v)
	if encodeErr != nil {
		panic("Failed to encode response")
	}
}

func (p *Handler) Error(w http.ResponseWriter, apiError err.ApiError) {
	p.AddHeader(w, "Content-Type", "application/json")
	w.WriteHeader(apiError.HttpStatusCode)
	encodeErr := json.NewEncoder(w).Encode(apiError)
	if encodeErr != nil {
		panic("Failed to encode response")
	}
}

func (p *Handler) AddHeader(w http.ResponseWriter, key, value string) {
	w.Header().Add(key, value)
}
