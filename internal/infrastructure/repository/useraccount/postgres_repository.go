package useraccount

import (
	"context"
	"database/sql"
	"errors"

	"jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/model/users"
)

var _ useraccount.UserRepository = (*PostgresUserRepository)(nil)

// PostgresUserRepository 实现 domain.UserRepository，属于基础设施层，委托 internal/model/users
type PostgresUserRepository struct {
	model users.UsersModel
}

// NewPostgresUserRepository 创建基于 PostgreSQL 的用户仓储
func NewPostgresUserRepository(model users.UsersModel) useraccount.UserRepository {
	return &PostgresUserRepository{model: model}
}

func (r *PostgresUserRepository) FindByPhone(ctx context.Context, phone string) (*useraccount.User, error) {
	u, err := r.model.FindOneByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return nil, useraccount.ErrUserNotFound
		}
		return nil, err
	}
	return toDomain(u), nil
}

func (r *PostgresUserRepository) FindByUserId(ctx context.Context, userId string) (*useraccount.User, error) {
	u, err := r.model.FindOneByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return nil, useraccount.ErrUserNotFound
		}
		return nil, err
	}
	return toDomain(u), nil
}

func (r *PostgresUserRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	return r.model.Contains(ctx, phone)
}

func (r *PostgresUserRepository) Add(ctx context.Context, user *useraccount.User) error {
	_, err := r.model.Insert(ctx, toPersistence(user))
	return err
}

func toDomain(u *users.Users) *useraccount.User {
	if u == nil {
		return nil
	}
	return &useraccount.User{
		UserId:    u.UserId,
		Phone:     u.Phone,
		Username:  u.Username,
		Password:  u.Password,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func toPersistence(u *useraccount.User) *users.Users {
	if u == nil {
		return nil
	}
	return &users.Users{
		UserId:    u.UserId,
		Phone:     u.Phone,
		Username:  u.Username,
		Password:  u.Password,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Status:    u.Status,
		DeletedAt: sql.NullTime{},
		// Id, CreatedAt, UpdatedAt 由 DB 或 go-zero 缓存层处理
	}
}
