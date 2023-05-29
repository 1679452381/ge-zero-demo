package dao

import (
	"context"
	"fmt"
	"go-zero-hello-2/mall/user/database"
	"go-zero-hello-2/mall/user/model"
)

type UserDao struct {
	*database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values (?,?)", user.TableName())
	_, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	return nil
}
