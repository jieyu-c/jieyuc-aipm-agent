package users

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		Contains(ctx context.Context, phone string) (bool, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (m *customUsersModel) Contains(ctx context.Context, phone string) (bool, error) {
	querySQL := fmt.Sprintf("select count(*) from %s where phone = $1", m.tableName())
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, querySQL, phone)
	return count > 0, err
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
