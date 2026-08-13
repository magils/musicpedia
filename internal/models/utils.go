package models

import (
	"context"
	"time"
)

func CreateContextWithTimeout(seconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
}

func CreateDefaultContext() (context.Context, context.CancelFunc) {
	return CreateContextWithTimeout(3)
}
