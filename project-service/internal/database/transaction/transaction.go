package transaction

import (
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/internal/database/gorms"
)

type Transaction struct {
	conn database.DbConn
}

func NewTransaction() *Transaction {
	return &Transaction{
		conn: gorms.NewTransaction(),
	}
}

//func (t *Transaction) Action(f func(conn database.DbConn) error) error {
//	t.conn.Begin()
//	err := f(t.conn)
//	var bErr *errs.BError
//	if errors.Is(err, bErr) {
//		bErr = err.(*errs.BError)
//		if bErr != nil {
//			t.conn.Rollback()
//			return bErr
//		} else {
//			t.conn.Commit()
//			return nil
//		}
//	}
//	if err != nil {
//		t.conn.Rollback()
//		return err
//	}
//	t.conn.Commit()
//	return nil
//}
