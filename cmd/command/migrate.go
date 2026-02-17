package command

import (
	"context"
	"sikabiz/user-importer/internal/config"
	"sikabiz/user-importer/internal/infra"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type MigrateCommand struct {
	Logger *log.Logger
}

func (cmd MigrateCommand) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "run migration",
		Run: func(_ *cobra.Command, args []string) {
			cmd.main(cfg, ctx, args)
		},
	}
}

func (cmd MigrateCommand) main(cfg *config.Config, ctx context.Context, args []string) {
	if len(args) == 0 {
		cmd.Logger.WithContext(ctx).Fatal("please specify migration command")
		return
	}

	psql, err := infra.NewPostgresClient(ctx, cfg.Database.Postgres)
	if err != nil {
		cmd.Logger.WithContext(ctx).Fatal(errors.Wrap(err, "server : failed to connect to postgresql"))
		return
	}

	migrationCommand := args[0]
	switch migrationCommand {
	case "up":
		err = psql.MigrateUp(cfg.Database.Postgres.Database)
		if err != nil {
			cmd.Logger.WithContext(ctx).Fatal(err)
			return
		}

	case "down":
		err = psql.MigrateDown(cfg.Database.Postgres.Database)
		if err != nil {
			cmd.Logger.WithContext(ctx).Fatal(err)
			return
		}

	default:
		cmd.Logger.WithContext(ctx).Fatal(errors.Errorf("migration command : %s is not supported", migrationCommand))
	}
}
