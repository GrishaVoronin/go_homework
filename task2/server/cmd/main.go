package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os/signal"
	"server/handlers"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", handlers.VersionHandler)
	mux.HandleFunc("/decode", handlers.DecodeHandler)
	mux.HandleFunc("/hard-op", handlers.HardOp)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("Starting server on port 8080")
		err := httpServer.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		fmt.Println("Start to shut down server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		fmt.Println("Server shutted down")
		return nil
	})

	err := group.Wait()
	if err != nil {
		fmt.Printf("after wait: %s\n", err)
		return
	}
}
