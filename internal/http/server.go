package http

import (
	"context"
	"net/http"

	peasant "github.com/candango/gopeasant"
	"github.com/candango/gopeasant/dummy"
	"github.com/candango/httpok"
	middleware "github.com/candango/httpok/middleware"
)

type NoncedServer struct {
	*httpok.GracefulServer
	peasant.NonceService
}

func NewNoncedServer(ctx context.Context) *NoncedServer {
	service := dummy.NewDummyInMemoryNonceService()
	s := &NoncedServer{
		GracefulServer: &httpok.GracefulServer{
			Name: "Cafofo Bastion",
			Server: &http.Server{
				Addr: ":8080",
			},
			Context: ctx,
			Logger:  &logrusRunLogger{},
		},
		NonceService: service,
	}
	s.buildHandler()
	return s
}

func (s *NoncedServer) buildHandler() {
	mux := http.NewServeMux()
	mux.Handle("/directory/", http.StripPrefix("/directory", NewDirectoryHandler()))
	mux.Handle("/new-nonce/", http.StripPrefix("/new-nonce", NewNonceHandler()))
	mux.Handle("/security/", http.StripPrefix("/security", peasant.Nonced(NewSecurityHandler(), s.NonceService)))
	s.Handler = middleware.Chain(mux,
		middleware.Logging(s.Logger),
		peasant.NonceServed(s.NonceService, ""),
	)
}
