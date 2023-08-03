package main

import (
	"cinema/src/comment-srv/db"
	"cinema/src/comment-srv/handler"
	"cinema/src/share/config"
	"cinema/src/share/pb"
	"cinema/src/share/utils/log"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.uber.org/zap"
)

func main() {

	log.Init("comment")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameComment),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is start ..."))
			// 注册redis
			//redisPool := share.NewRedisPool(3, 3, 1,300*time.Second,":6379","redis")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterCommentServiceExtHandler(service.Server(), handler.NewCommentServiceExtHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("comment-srv服务启动失败 ...")
	}
}
