package utils

import "context"

type ContextKey struct {
	name string
}

var ContextKeyUserId = &ContextKey{name: "userId"}

func (c *ContextKey) GetStringValueFromContext(ctx context.Context) string {
	return ctx.Value(c).(string)
}
