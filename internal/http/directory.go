package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/candango/httpok/middleware"
)

type DirectoryHandler struct {
	*http.ServeMux
}

func NewDirectoryHandler() http.Handler {
	h := &DirectoryHandler{
		&http.ServeMux{},
	}
	h.HandleFunc("GET /", h.getDirectory)
	return middleware.ExactPath("/", h)
}

func (h *DirectoryHandler) getDirectory(w http.ResponseWriter, r *http.Request) {
	// TODO: check the protocol from the header
	proto := "http"
	if r.TLS != nil {
		proto = "https"
	}
	currentHost := fmt.Sprintf("%s://%s", proto, r.Host)

	directory := map[string]any{
		"new-nonce": fmt.Sprintf("%s/new-nonce/", currentHost),
		"security": map[string]string{
			"auth": fmt.Sprintf("%s/security/auth/", currentHost),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(directory)
	if err != nil {
		http.Error(w, fmt.Sprintf("error build response, %v", err),
			http.StatusInternalServerError)
		return
	}
}
