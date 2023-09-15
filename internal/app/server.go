package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/http-loki/internal/app/mux"
	"github.com/g8rswimmer/http-loki/internal/config"
)

func Run(values *config.Values) error {

	srvr := http.Server{
		Addr:    ":" + values.Port,
		Handler: mux.Handler(values),
	}

	go func() {
		err := srvr.ListenAndServe()
		switch {
		case errors.Is(err, http.ErrServerClosed):
		case err != nil:
			log.Panicf("server error %v", err)
		default:
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srvr.Shutdown(ctx)
}
