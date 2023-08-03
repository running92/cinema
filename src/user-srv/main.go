package main

import (
	"cinema/src/share/config"
	"cinema/src/share/pb"
	"cinema/src/share/utils/log"
	"cinema/src/user-srv/db"
	"cinema/src/user-srv/handler"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"

	"go.uber.org/zap"
)

func main() {

	log.Init("user")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameUser),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))
			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceExtHandler(service.Server(), handler.NewUserServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败 ...")
	}
}
