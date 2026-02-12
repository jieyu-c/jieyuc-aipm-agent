package useraccount

import "context"

// UserRepository 用户仓储接口，由应用层依赖，由基础设施层实现
type UserRepository interface {
	FindByPhone(ctx context.Context, phone string) (*User, error)
	FindByUserId(ctx context.Context, userId string) (*User, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	Add(ctx context.Context, user *User) error
}
