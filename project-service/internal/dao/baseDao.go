package dao

import (
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-service/internal/database/gorms"
)

type baseDao struct {
	conn *gorms.GormConn
}

func (b *baseDao) getConn(tx database.DbConn) *gorms.GormConn {
	var conn *gorms.GormConn
	if tx != nil {
		conn = tx.(*gorms.GormConn)
	} else {
		conn = b.conn
	}
	return conn
}
