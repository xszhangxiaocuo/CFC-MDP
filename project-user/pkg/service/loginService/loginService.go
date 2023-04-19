package loginService

import (
	"context"
	common "github.com/xszhangxiaocuo/CFC-MDP/project-common"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	loginRpc "github.com/xszhangxiaocuo/CFC-MDP/project-grpc/user/login"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database/transaction"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/domain"
	"log"
)

type LoginService struct {
	loginRpc.UnimplementedLoginServiceServer
	accountDomain *domain.AccountDomain
	tx            *transaction.Transaction
}

func NewLoginService() *LoginService {
	return &LoginService{
		accountDomain: domain.NewAccountDomain(),
		tx:            transaction.NewTransaction(),
	}
}

func (ls *LoginService) Login(ctx context.Context, req *loginRpc.LoginRequest) (*loginRpc.LoginResponse, error) {
	// 1. 查询账号是否存在
	user, err := ls.accountDomain.FindAccount(nil, req.Account)

	if err != nil {
		log.Println("account search failed:", err)
		return nil, err
	}
	if user == nil {
		return &loginRpc.LoginResponse{Result: "unexist"}, nil
	}
	//2.查询密码是否正确
	if req.Password != user.Password {
		return nil, errors.NewError(common.PASSWDWRONG, "password wrong")
	}
	return &loginRpc.LoginResponse{Result: "success"}, nil
}
