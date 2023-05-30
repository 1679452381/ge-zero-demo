package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-hello-2/mall/user/database"
	"go-zero-hello-2/mall/user/model"
	"strconv"
)

type UserDao struct {
	*database.DBConn
}

var cacheUserIdPrefix = "cache:user:id:"

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values (?,?)", user.TableName())
	res, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = strconv.Itoa(int(id))
	return nil
}

func (d *UserDao) FindOneById(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	fmt.Println("dao")
	fmt.Println(user)
	query := fmt.Sprintf("select * from %s where id=?  limit 1", user.TableName())
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	//使用cache连接
	err := d.CacheConn.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}
