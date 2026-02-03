package utils

import (
	"regexp"
	"unicode"
)

var (
	// 中国大陆手机号正则：1开头，第二位为3-9，后面9位数字
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
	// 验证码正则：6位数字
	verifyCodeRegex = regexp.MustCompile(`^\d{6}$`)
)

// ValidatePhone 校验手机号格式
// 返回 true 表示格式正确
func ValidatePhone(phone string) bool {
	if phone == "" {
		return false
	}
	return phoneRegex.MatchString(phone)
}

// ValidateVerifyCode 校验验证码格式（6位数字）
func ValidateVerifyCode(code string) bool {
	if code == "" {
		return false
	}
	return verifyCodeRegex.MatchString(code)
}

// PasswordStrength 密码强度等级
type PasswordStrength int

const (
	PasswordWeak     PasswordStrength = iota // 弱：不满足基本要求
	PasswordMedium                           // 中：满足基本要求
	PasswordStrong                           // 强：包含多种字符类型
)

// PasswordValidationResult 密码校验结果
type PasswordValidationResult struct {
	Valid    bool             // 是否有效
	Strength PasswordStrength // 强度等级
	Message  string           // 提示信息
}

// ValidatePassword 校验密码强度
// 要求：
// - 长度 8-32 位
// - 至少包含数字和字母
// - 建议包含特殊字符
func ValidatePassword(password string) PasswordValidationResult {
	length := len(password)

	// 长度校验
	if length < 8 {
		return PasswordValidationResult{
			Valid:    false,
			Strength: PasswordWeak,
			Message:  "密码长度不能少于8位",
		}
	}
	if length > 32 {
		return PasswordValidationResult{
			Valid:    false,
			Strength: PasswordWeak,
			Message:  "密码长度不能超过32位",
		}
	}

	var (
		hasLower   bool
		hasUpper   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	hasLetter := hasLower || hasUpper

	// 必须同时包含字母和数字
	if !hasLetter || !hasDigit {
		return PasswordValidationResult{
			Valid:    false,
			Strength: PasswordWeak,
			Message:  "密码必须同时包含字母和数字",
		}
	}

	// 计算强度
	strength := PasswordMedium
	if hasLower && hasUpper && hasDigit && hasSpecial {
		strength = PasswordStrong
	}

	return PasswordValidationResult{
		Valid:    true,
		Strength: strength,
		Message:  "",
	}
}
