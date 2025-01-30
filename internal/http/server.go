package http

import (
	"context"
	"net/http"

	"github.com/candango/httpok"
)

type NoncedServer struct {
	*httpok.GracefulServer
	Count int
}

func NewNoncedServer(ctx context.Context) *NoncedServer {
	s := &NoncedServer{
		GracefulServer: &httpok.GracefulServer{
			Name: "Cafofo Bastion",
			Server: &http.Server{
				Addr: ":8080",
			},
			Context:   ctx,
			RunLogger: &logrusRunLogger{},
		},
		Count: 0,
	}
	s.buildHandler()
	return s
}

func (s *NoncedServer) buildHandler() {
	mux := http.NewServeMux()
	mux.Handle("/directory/", http.StripPrefix("/directory", NewDirectoryHandler()))
	mux.Handle("/new-nonce/", http.StripPrefix("/new-nonce", NewNonceHandler()))
	s.Handler = httpok.Chain(mux)
}
