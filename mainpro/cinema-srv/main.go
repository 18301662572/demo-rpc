package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"my-micro/demo/src/share/config"
	"my-micro/demo/src/share/pb"
	"my-micro/demo/src/share/utils/log"
	"my-micro/demo/src/cinema-srv/handler"
	"my-micro/demo/src/cinema-srv/db"
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
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("cinema-srv", "cinema-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterCinemaServiceExtHandler(service.Server(), handler.NewCinemaServiceExtHandler(), server.InternalHandler(true))
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
