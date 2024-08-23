package postgres

import (
	pb "gitlab.com/lingualeap/auth/genprotos/users"
)

type InitRoot interface {
	User() User
}

type User interface {
	Register(user *pb.UserReq) (*pb.Void, error)
	Login(user *pb.UserLogin) (*pb.User, error)
	GetById(id *pb.ById) (*pb.User, error)
	GetAll(filter *pb.UserFilter) (*pb.AllUsers, error)
	Update(user *pb.User) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	ChangePassword(changePass *pb.ChangePass) (*pb.Void, error)
	ForgotPassword(forgotPass *pb.ForgotPass) (*pb.Void, error)
	ResetPassword(resetPass *pb.ResetPass) (*pb.Void, error)
}
