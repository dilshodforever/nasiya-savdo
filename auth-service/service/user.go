package service

import (
	"context"
	"log"

	pb "gitlab.com/lingualeap/auth/genprotos/users"
	s "gitlab.com/lingualeap/auth/storage"
)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) Register(ctx context.Context, user *pb.UserReq) (*pb.Void, error) {
	void, err := c.stg.User().Register(user)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) Login(ctx context.Context, login *pb.UserLogin) (*pb.User, error) {
	user, err := c.stg.User().Login(login)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
	user, err := c.stg.User().GetById(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (c *UserService) GetAll(ctx context.Context, filter *pb.UserFilter) (*pb.AllUsers, error) {
	users, err := c.stg.User().GetAll(filter)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}

func (c *UserService) Update(ctx context.Context, user *pb.User) (*pb.Void, error) {
	void, err := c.stg.User().Update(user)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	void, err := c.stg.User().Delete(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ChangePassword(ctx context.Context, changePass *pb.ChangePass) (*pb.Void, error) {
	void, err := c.stg.User().ChangePassword(changePass)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ForgotPassword(ctx context.Context, forgotPass *pb.ForgotPass) (*pb.Void, error) {
	void, err := c.stg.User().ForgotPassword(forgotPass)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}

func (c *UserService) ResetPassword(ctx context.Context, resetPass *pb.ResetPass) (*pb.Void, error) {
	void, err := c.stg.User().ResetPassword(resetPass)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return void, nil
}
