package main

import (
	"cinema/src/cms-srv/db"
	"cinema/src/cms-srv/handler"
	"cinema/src/share/config"
	"cinema/src/share/pb"
	"cinema/src/share/utils/log"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.uber.org/zap"
)

func main() {

	log.Init("cms")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCMS),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCMSServiceExtHandler(service.Server(), handler.NewCMSServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("cms-srv", "cms-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("cms-srv服务启动失败 ...")
	}
}
