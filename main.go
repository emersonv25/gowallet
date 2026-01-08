package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emersonv25/gowallet/config"
	"github.com/emersonv25/gowallet/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("Failed to initialize configuration: %s", err.Error())
		return
	}

	router.Initialize()
	server := &http.Server{Addr: ":9000"}

	// Canal de sinais
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start do servidor
	go func() {
		logger.Info("Servidor iniciado na porta 9000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Erro no servidor: %v", err)
		}
	}()

	// Aguarda Ctrl+C
	<-stop
	logger.Info("Recebido sinal de encerramento")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Erro ao desligar servidor: %v", err)
	}

	logger.Info("Servidor finalizado corretamente")
}
