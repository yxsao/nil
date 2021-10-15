package main

import (
	"github.com/yxsao/nil/tree/master/newmicro/user/handler"
	pb "github.com/yxsao/nil/tree/master/newmicro/user/proto"

	"github.com/micro/micro/v2/service"
	"github.com/micro/micro/v2/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
