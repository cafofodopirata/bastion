package http

import (
	"net/http"

	peasant "github.com/candango/gopeasant"
)

type NonceHandler struct {
	*http.ServeMux
}

func NewNonceHandler() *NonceHandler {
	h := &NonceHandler{
		&http.ServeMux{},
	}
	h.HandleFunc("GET /", h.newNonce)
	return h
}

func (h *NonceHandler) newNonce(w http.ResponseWriter, r *http.Request) {
	log := &logrusRunLogger{}
	s, ok := r.Context().Value("nonce-service").(peasant.NonceService)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("nonce service not present in the context")
		return
	}
	nonce, err := s.GetNonce(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("error getting a new nonce %v", err)
		return
	}
	w.Header().Add("nonce", nonce)
}
