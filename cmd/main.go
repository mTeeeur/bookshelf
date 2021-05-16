package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/mTeeeur/bookshelf"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		call := <-ch
		log.Println(call)
		cancel()
	}()

	r := mux.NewRouter()
	s := bookshelf.NewServer("127.0.0.1:5000", r)

	if err := s.Serve(ctx); err != nil {
		log.Fatal(err)
	}
}
