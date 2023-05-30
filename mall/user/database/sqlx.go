package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	CacheConn sqlc.CachedConn
}

func ConnectDb(dataSource string, conf cache.CacheConf) *DBConn {
	conn := sqlx.NewMysql(dataSource)
	d := &DBConn{
		Conn: conn,
	}
	if conf != nil {
		cacheConn := sqlc.NewConn(conn, conf)
		d.CacheConn = cacheConn
	}
	return d
}
