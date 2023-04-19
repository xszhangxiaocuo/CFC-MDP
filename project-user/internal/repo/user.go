package repo

import (
	"context"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/model"
)

type UserRepo interface {
	FindUserByAccount(ctx context.Context, tx database.DbConn, account string) (*model.User, *errors.Error)
}
