package http

import (
	"net/http"
)

type SecurityHandler struct {
	*http.ServeMux
}

func NewSecurityHandler() http.Handler {
	h := &SecurityHandler{
		&http.ServeMux{},
	}
	h.HandleFunc("POST /auth/", h.auth)
	return h
}

func (h *SecurityHandler) auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tamo autenticando"))
}
