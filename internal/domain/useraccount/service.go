package useraccount

// DomainService 用户账户领域服务，仅包含纯领域规则，无 I/O
type DomainService struct{}

// CanLogin 判断用户是否允许登录（如未禁用），不允许则返回领域错误
func (DomainService) CanLogin(user *User) error {
	if user == nil {
		return ErrUserNotFound
	}
	if user.Status == UserStatusBanned {
		return ErrUserBanned
	}
	return nil
}
