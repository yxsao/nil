package main

import (
	"context"
	"fmt"
	capMicro "ysmod/proto/cap"

	"github.com/micro/go-micro/v2"
)

type YsServer struct { //需要实现的方法
}

func (c *YsServer) SayHello(ctx context.Context, req *capMicro.SayInfo, res *capMicro.ReInfo) error {
	res.Rmess = "收到" + req.Mess
	return nil
}

func main() {
	//创建新的服务
	ysSer := micro.NewService(
		micro.Name("ys.server"),
	)
	ysSer.Init() //初始化方法
	//注册服务
	err := capMicro.RegisterYsServerHandler(ysSer.Server(), new(YsServer))
	if err != nil {
		return
	}
	//运行服务
	if err := ysSer.Run(); err != nil {
		fmt.Println(err)
	}
}
