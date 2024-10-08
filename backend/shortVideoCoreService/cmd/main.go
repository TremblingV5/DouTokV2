package main

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	c := &conf.Config{}
	launcher.New(
		launcher.WithConfigValue(c),
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
		launcher.WithGrpcServer(func(configValue interface{}) *grpc.Server {
			cfg, ok := configValue.(*conf.Config)
			if !ok {
				panic("invalid config value")
			}
			// init global resources
			utils.InitDefaultSnowflakeNode(cfg.App.Node)
			query.SetDefault(mysqlx.GetDBClient(context.Background()))
			logger := utils.InitStdLogger(&cfg.App)

			return server.NewGRPCServer(cfg, logger)
		}),
	).Run()
}
