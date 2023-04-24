package repo

import (
	"context"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/common/gen/common/gen/model"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/errors"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database"
)

type TemplateRepo interface {
	SaveTemplate(ctx context.Context, tx database.DbConn, tem *model.Template) (bool, *errors.Error)
	GetTemplate(ctx context.Context, tx database.DbConn, id int32) (*model.Template, *errors.Error)
}
