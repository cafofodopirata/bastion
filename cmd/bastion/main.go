package main

import (
	"context"
	"os"

	"github.com/cafofodopirata/bastion/internal/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	s := http.NewNoncedServer(ctx)
	<-s.Run()
	cancel()
	defer os.Exit(0)
	return
}
