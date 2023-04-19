package dao

import (
	"context"
	common "github.com/xszhangxiaocuo/CFC-MDP/project-common"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database/gorms"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/model"
	"gorm.io/gorm"
)

type UserDao struct {
	baseDao
}

func NewUserDao() *UserDao {
	return &UserDao{baseDao{conn: gorms.NewTransaction()}}
}

func (u *UserDao) FindUserByAccount(ctx context.Context, tx database.DbConn, account string) (*model.User, *errors.Error) {
	var user *model.User
	conn := u.conn
	err := conn.Session(ctx).Where("account = ? and deleted=0", account).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewError(common.ACCOUNTNOTFIND, "account absent")
	} else if err != nil {
		return nil, errors.NewError(common.FINDERR, err.Error())
	}
	return user, nil
}
