package handler

import (
	"context"
	go_micro_service_user "github.com/yxsao/nil/tree/master/newmicro/user/proto/user"

	"github.com/yxsao/nil/tree/master/newmicro/user/domain/model"
	"github.com/yxsao/nil/tree/master/newmicro/user/domain/service"
)

type User struct {
	UserDataService service.IuserDataService
}

// Register 注册
func (u *User) Register(ctx context.Context, userRegisterRequest *go_micro_service_user.UserRegisterRequest,
	userRegisterResponse *go_micro_service_user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

// Login 登录
func (u *User) Login(ctx context.Context, userLogin *go_micro_service_user.UserLoginRequest,
	loginResponse *go_micro_service_user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return err
	}
	loginResponse.IsSuccess = isOk
	return nil
}

// GetUserInfo 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *go_micro_service_user.UserInfoRequest,
	userInfoResponse *go_micro_service_user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserForResponse(userInfo)
	return nil
}

// UserForResponse 类型转化
func UserForResponse(userModel *model.User) *go_micro_service_user.UserInfoResponse {
	response := &go_micro_service_user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
