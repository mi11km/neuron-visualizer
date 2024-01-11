package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/mi11km/neuron-visualizer/server/interfaces"
	"github.com/mi11km/neuron-visualizer/server/openapi"
)

type config struct {
	Port                int    `env:"PORT" default:"8080"`
	NeuronSimulationDir string `env:"NEURON_SIMULATION_DIR" default:"./simulations"`
}

func main() {
	cfg := config{}
	if err := envconfig.Process("", &cfg); err != nil {
		slog.Error("envconfig.Process", slog.Any("err", err))
		os.Exit(1)
	}

	// Logging
	logger := httplog.NewLogger(
		"", httplog.Options{
			JSON:           true,
			LogLevel:       slog.LevelDebug,
			RequestHeaders: true,
		},
	)
	slog.SetDefault(logger.Logger)

	// Server
	neuronVisualizer, err := interfaces.NewNeuronVisualizerServer(cfg.NeuronSimulationDir)
	if err != nil {
		slog.Error("interfaces.NewNeuronVisualizerHandler", slog.Any("err", err))
		os.Exit(1)
	}

	mux := openapi.HandlerWithOptions(
		neuronVisualizer, openapi.ChiServerOptions{
			Middlewares: []openapi.MiddlewareFunc{
				middleware.Recoverer,
				httplog.RequestLogger(logger),
			},
		},
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		slog.Info("starting server", slog.Any("port", cfg.Port))
		err := server.ListenAndServe()
		if err == nil {
			return
		}
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("server closed")
			return
		}
		errCh <- err
	}()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		slog.Info("shutting down server...")
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("server.Shutdown", slog.Any("err", err))
		}
	}()

	signalCh := make(chan os.Signal, 1)
	defer close(signalCh)
	signal.Notify(signalCh, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case err := <-errCh:
		slog.Error("server terminated with error", slog.Any("err", err))
	case sig := <-signalCh:
		slog.Info("server terminated with signal", slog.Any("signal", sig.String()))
	}
}
