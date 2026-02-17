package command

import (
	"context"
	"encoding/json"
	"os"
	"sikabiz/user-importer/internal/config"
	"sikabiz/user-importer/internal/domain"
	"sikabiz/user-importer/internal/infra"
	"sikabiz/user-importer/internal/repository"
	userService "sikabiz/user-importer/internal/service/user"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ImportCommand struct {
	Logger *log.Logger
}

func (cmd ImportCommand) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "import",
		Short: "run importer",
		Run: func(_ *cobra.Command, args []string) {
			cmd.main(cfg, ctx, args)
		},
	}
}

func (cmd ImportCommand) main(cfg *config.Config, ctx context.Context, args []string) {
	psql, err := infra.NewPostgresClient(ctx, cfg.Database.Postgres)
	if err != nil {
		cmd.Logger.WithContext(ctx).Fatal(errors.Wrap(err, "importer : failed to connect to postgresql"))
		return
	}

	file, err := os.Open("users_data.json")
	if err != nil {
		cmd.Logger.WithContext(ctx).Fatal(errors.Wrap(err, "importer : failed to open users data"))
		return
	}

	defer file.Close()

	var users []domain.User

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		cmd.Logger.WithContext(ctx).Fatal(errors.Wrap(err, "importer : failed to decode users data"))
		return
	}

	userRepo := repository.NewUserRepository(psql.GetDb())
	addressRepo := repository.NewAddressRepository(psql.GetDb())

	userServiceInstance := userService.NewUserService(userRepo, addressRepo, cmd.Logger)

	if errs := userServiceInstance.ImportUsers(ctx, users, cfg.WorkerCount); err != nil {
		for _, err := range errs {
			cmd.Logger.WithContext(ctx).Error(errors.Wrap(err, "importer : failed to insert users data"))
		}
	}
}
