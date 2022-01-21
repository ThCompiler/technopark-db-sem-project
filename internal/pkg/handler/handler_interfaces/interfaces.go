package handler_interfaces

import (
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(ctx echo.Context) error

func (f HandlerFunc) ServeHTTP(ctx echo.Context) error {
	return f(ctx)
}

type Handler interface {
	ServeHTTP(ctx echo.Context) error
}
type HMiddlewareFunc func(Handler) Handler
type HFMiddlewareFunc func(HandlerFunc) HandlerFunc
