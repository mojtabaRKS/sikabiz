package command

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"sikabiz/user-importer/internal/api"
	userHandler "sikabiz/user-importer/internal/api/handler/user"
	"sikabiz/user-importer/internal/config"
	"sikabiz/user-importer/internal/infra"
	"sikabiz/user-importer/internal/repository"
	userService "sikabiz/user-importer/internal/service/user"
)

type Server struct {
	Logger *logrus.Logger
}

func (cmd Server) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "run Gateway server",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(cfg, ctx)
		},
	}
}

func (cmd Server) main(cfg *config.Config, ctx context.Context) {
	psql, err := infra.NewPostgresClient(ctx, cfg.Database.Postgres)
	if err != nil {
		cmd.Logger.WithContext(ctx).Fatal(errors.Wrap(err, "server : failed to connect to postgresql"))
		return
	}

	userRepository := repository.NewUserRepository(psql.GetDb())
	addressRepository := repository.NewAddressRepository(psql.GetDb())

	userServiceInstance := userService.NewUserService(userRepository, addressRepository, cmd.Logger)

	userHandlerInstance := userHandler.NewUserHandler(userServiceInstance)

	server := api.New(cfg.AppEnv)
	server.SetupAPIRoutes(
		userHandlerInstance,
	)

	if err := server.Serve(ctx, fmt.Sprintf(":%d", cfg.HTTP.Port)); err != nil {
		cmd.Logger.Fatal(err)
	}
}
