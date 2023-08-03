package main

import (
	"cinema/src/cinema-srv/db"
	"cinema/src/cinema-srv/handler"
	"cinema/src/share/config"
	"cinema/src/share/pb"
	"cinema/src/share/utils/log"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.uber.org/zap"
)

func main() {

	log.Init("cinema")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCinema),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("cinema-srv", "cinema-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCinemaServiceExtHandler(service.Server(), handler.NewCinemaServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("cinema-srv", "cinema-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("cinema-srv服务启动失败 ...")
	}
}
