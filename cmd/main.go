package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"thiagodotjpeg/pulse/internal/database"
	"thiagodotjpeg/pulse/internal/monitor/controller"
	"time"
)

func main() {
	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	errorServ := make(chan error, 1)

	slog.SetDefault(slogLogger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := database.NewPool(ctx, "postgresql://teste:teste@localhost:5432/teste")
	if err != nil {
		slogLogger.Error("Ocorreu um erro ao criar o pool de conexões", "error", err)
		os.Exit(1)
	}

	handler := controller.NewRouter(slogLogger, pool)
	
	srv := &http.Server{
		Addr: ":8000",
		Handler: handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		slogLogger.Info("Servidor HTTP escutando na", "port", 8000)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorServ <- err
			return
		}
		errorServ <- nil
	}()

	select {
	case err := <-errorServ:
		slogLogger.Error("Falha ao iniciar o servidor", "error", err)
		os.Exit(1)
	case <-ctx.Done():
		slogLogger.Info("Sinal de encerramento recebido, desligando servidor...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil{
		os.Exit(1)
	}

	pool.Close()
	slogLogger.Info("Servidor encerrado com sucesso")
}