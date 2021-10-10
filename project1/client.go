package main

import (
	"context"
	"fmt"
	cap "ysmod/proto/cap"

	"github.com/micro/go-micro/v2"
)

func main() {
	//实例化
	ysClient := micro.NewService(
		micro.Name("ys.client"),
	)
	//初始化
	ysClient.Init()

	Isao := cap.NewYsServerService("ys.server", ysClient.Client())
	res, err := Isao.SayHello(context.TODO(), &cap.SayInfo{Mess: "123"})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Rmess)
}
