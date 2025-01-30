package http

import (
	"net/http"

	"math/rand"
)

func randomString(s int) string {
	asciiLower := "abcdefghijklmnopqrstuvwxyz"
	asciiUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "012345679"
	chars := []rune(asciiLower + asciiUpper + digits)
	r := make([]rune, s)
	for i := range r {
		r[i] = chars[rand.Intn(len(chars))]
	}
	return string(r)
}

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
	w.Header().Add("nonce", randomString(45))
}
