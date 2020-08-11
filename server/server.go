package server

import (
	"context"
	"github.com/gorilla/mux"
	"gostart/utils/conf"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func GracefullyShutdown(r *mux.Router) error {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()

	srv := &http.Server{
		Handler: r,
		Addr: conf.GetFullAddr(),
		WriteTimeout: conf.GetCtxTimeout(),
		ReadTimeout:  conf.GetCtxTimeout(),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+s\n", err)
		}
	}()
	log.Printf("listening on port %v", conf.GetPort())

	<-ctx.Done()
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), conf.GetCtxTimeout())
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatalf("server Shutdown Failed: %+s", err)
	}
	log.Printf("server exited properly")

	return err
}
