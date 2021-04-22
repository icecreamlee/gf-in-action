package service

import "context"

const ContextKey = "UserContext"

var Context = &userContext{}

type userContext struct {
	UserId   int
	Username string
}

func (c userContext) Get(ctx context.Context) *userContext {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*userContext); ok {
		return localCtx
	}
	return nil
}
