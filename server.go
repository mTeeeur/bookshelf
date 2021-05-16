package bookshelf

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string, h http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         addr,
			Handler:      h,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
		},
	}
}

func (s *Server) Serve(ctx context.Context) (err error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Println("server started")

	<-ctx.Done()

	log.Println("server stopped")

	shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}

	log.Println("server exited")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
