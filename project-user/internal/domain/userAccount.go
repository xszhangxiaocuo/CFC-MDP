package domain

import (
	"context"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/dao"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/model"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/repo"
	"time"
)

type AccountDomain struct {
	userAccountRepo repo.UserRepo
	//userRpcDomain   *UserRpcDomain
}

func NewAccountDomain() *AccountDomain {
	return &AccountDomain{
		userAccountRepo: dao.NewUserDao(),
		//userRpcDomain:   NewUserRpcDomain(),
	}
}

func (d *AccountDomain) FindAccount(conn database.DbConn, account string) (*model.User, *errors.Error) {
	user := &model.User{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	user, err := d.userAccountRepo.FindUserByAccount(c, conn, account)
	if err != nil {
		return nil, err
	}
	return user, nil
}
