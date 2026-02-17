package main

import (
	"context"
	"os/signal"
	"sikabiz/user-importer/cmd/command"
	"sikabiz/user-importer/internal/config"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	const description = "sikabiz user importer"
	root := &cobra.Command{Short: description}

	cfg, err := config.Load()
	if err != nil {
		log.WithContext(ctx).Fatal(err)
	}

	logger := log.New()

	root.AddCommand(
		command.Server{Logger: logger}.Command(ctx, cfg),
		command.MigrateCommand{Logger: logger}.Command(ctx, cfg),
		command.ImportCommand{Logger: logger}.Command(ctx, cfg),
	)

	if err := root.Execute(); err != nil {
		logger.WithContext(ctx).Fatalf("failed to execute root command: \n%v", err)
	}
}
