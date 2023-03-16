package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/santi1234567/eth-pools-identifier/config"
	"github.com/santi1234567/eth-pools-identifier/poolIdentifier"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.NewCliConfig()
	if err != nil {
		log.Fatal(err)
	}
	logLevel, err := log.ParseLevel(config.Verbosity)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(logLevel)

	poolIdentifier, err := poolIdentifier.NewPoolIdentifier(context.Background(), config)
	
	if err != nil {
		log.Fatal(err)
	}

	go poolIdentifier.Run()
	// Wait for signal.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	for {
		sig := <-sigCh
		if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == os.Interrupt || sig == os.Kill {
			break
		}
	}

	log.Info("Stopping eth-pools-identifier...")
}