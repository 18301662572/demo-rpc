package main

import (
	"code.oldbody.com/studygolang/demo(rpc)/mainpro/share/config"
	"code.oldbody.com/studygolang/demo(rpc)/mainpro/share/pb"
	"code.oldbody.com/studygolang/demo(rpc)/mainpro/share/utils/log"
	"code.oldbody.com/studygolang/demo(rpc)/mainpro/user-srv/db"
	"code.oldbody.com/studygolang/demo(rpc)/mainpro/user-srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
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
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))

			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceExtHandler(service.Server(), handler.NewUserServiceExtHandler(), server.InternalHandler(true))
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
