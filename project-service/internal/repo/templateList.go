package repo

import (
	"context"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/model"
)

type TemplateListRepo interface {
	GetTemplateList(ctx context.Context, tx database.DbConn, tem *model.TemplateList) (bool, *errors.Error)
}
