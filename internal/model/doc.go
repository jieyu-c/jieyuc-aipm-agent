// Package model 提供与数据库表一一对应的持久化模型及驱动注册。
//
// 在 DDD 中归属「基础设施层」：仅被 internal/infrastructure/repository 等
// 基础设施实现使用，领域层与应用层不依赖本包。
package model
