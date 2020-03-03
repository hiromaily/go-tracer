package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/opentracing/opentracing-go"
)

// Server is server object
type Server struct {
	port   int
	tracer opentracing.Tracer
}

// NewServer is to return server object
func NewServer(port int, trs opentracing.Tracer) *Server {
	return &Server{
		port:   port,
		tracer: trs,
	}
}

// Listen is to start listening
func (s *Server) Listen() {
	//TODO
	//https://go.googlesource.com/playground/+/0cf5f3c26a052bbbbb37ea3ee739d474b27c98bf/server.go

	stopCh := make(chan os.Signal)
	// nolint:staticcheck
	signal.Notify(stopCh, os.Interrupt)

	//server object

	http.HandleFunc("/favicon.ico", s.favicon())
	http.HandleFunc("/", s.handler())
	srv := &http.Server{Addr: fmt.Sprintf(":%d", s.port), Handler: http.DefaultServeMux}
	log.Printf("Server start with port %d ...", s.port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-stopCh // wait for SIGINT

	log.Println("Shutting down server...")
	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cancel()

	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}

func (s *Server) handler() func(http.ResponseWriter, *http.Request) {
	//TODO
	//https://opentracing.io/guides/golang/quick-start/

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Header, r.Host)

		//tracer := opentracing.GlobalTracer()
		msg := "say-hello"
		span := s.tracer.StartSpan(msg)
		defer span.Finish()

		span.SetTag("tag", "something")

		fmt.Fprintf(w, "Hello, %s!", msg)
	}
}

func (s *Server) favicon() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
