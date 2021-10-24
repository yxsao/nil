package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/yxsao/nil/tree/master/newmicro/user/domain/repository"
	"github.com/yxsao/nil/tree/master/newmicro/user/domain/service"
	"github.com/yxsao/nil/tree/master/newmicro/user/handler"
	go_micro_service_user "github.com/yxsao/nil/tree/master/newmicro/user/proto/user"
)

func main() {
	// 服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql",
		//"root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
		"root:aoxn@(127.0.0.1:3306)/teys?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)

	//只执行一次，数据表初始化
	//rp := repository.NewUserRepository(db)
	//err = rp.InitTable()
	//if err != nil {
	//	return
	//}

	//创建服务实例
	userDataService := service.NewUserDataService(repository.NewUserRepository(db))

	//注册Handler
	err = go_micro_service_user.RegisterUserHandler(srv.Server(),
		&handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
