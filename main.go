package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/YvanJAquino/meet-chatbot-101/env"
	"github.com/YvanJAquino/meet-chatbot-101/handlers"
)

func main() {
	env := env.NewManager()
	env.Get()

	parent := context.Background()
	notify, stop := signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.DefaultHandler)
	server := &http.Server{
		Addr:        ":" + env.Port,
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return parent },
	}
	fmt.Printf("Starting HTTP server on localhost:%s\n", env.Port)
	go server.ListenAndServe()
	<-notify.Done() // Allow closing of the server when context is done.
	shutCtx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	server.Shutdown(shutCtx)
}
