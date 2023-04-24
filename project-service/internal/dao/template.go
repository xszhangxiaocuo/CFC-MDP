package dao

import (
	"context"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database/gorms"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/model"
)

type TemplateDao struct {
	baseDao
}

func NewTemplateDao() *TemplateDao {
	return &TemplateDao{baseDao{gorms.NewTransaction()}}
}

func (t *TemplateDao) SaveTemplate(ctx context.Context, tx database.DbConn, tem *model.Template) (bool, *errors.Error) {

}
