package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

// 局部中间件
func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("example middle")
	return func(w http.ResponseWriter, r *http.Request) {
		// Passthrough to next handler if need
		next(w, r)
	}
}
