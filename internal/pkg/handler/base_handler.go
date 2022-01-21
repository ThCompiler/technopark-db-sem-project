package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	hf "tech-db-forum/internal/pkg/handler/handler_interfaces"
	"tech-db-forum/internal/pkg/utilits"
	"tech-db-forum/internal/pkg/utilits/delivery"

	"github.com/sirupsen/logrus"
)

const (
	GET     = http.MethodGet
	POST    = http.MethodPost
	PUT     = http.MethodPut
	DELETE  = http.MethodDelete
	OPTIONS = http.MethodOptions
)

type BaseHandler struct {
	handlerMethods map[string]hf.HandlerFunc
	middlewares    []hf.HMiddlewareFunc
	HelpHandlers
}

func NewBaseHandler(log *logrus.Logger) *BaseHandler {
	h := &BaseHandler{handlerMethods: map[string]hf.HandlerFunc{}, middlewares: []hf.HMiddlewareFunc{},
		HelpHandlers: HelpHandlers{
			ErrorConvertor: delivery.ErrorConvertor{
				Responder: delivery.Responder{
					LogObject: utilits.NewLogObject(log),
				},
			},
		},
	}
	return h
}

func (h *BaseHandler) AddMiddleware(middleware ...hf.HMiddlewareFunc) {
	h.middlewares = append(h.middlewares, middleware...)
}

func (h *BaseHandler) AddMethod(method string, handlerMethod hf.HandlerFunc, middlewares ...hf.HFMiddlewareFunc) {
	h.handlerMethods[method] = h.applyHFMiddleware(handlerMethod, middlewares...)
}

func (h *BaseHandler) applyHFMiddleware(handlerMethod hf.HandlerFunc,
	middlewares ...hf.HFMiddlewareFunc) hf.HandlerFunc {
	resultHandlerMethod := handlerMethod
	for index := len(middlewares) - 1; index >= 0; index-- {
		resultHandlerMethod = middlewares[index](resultHandlerMethod)
	}
	return resultHandlerMethod
}

func (h *BaseHandler) applyMiddleware(handler hf.Handler) echo.HandlerFunc {
	resultHandler := handler
	for index := len(h.middlewares) - 1; index >= 0; index-- {
		resultHandler = h.middlewares[index](resultHandler)
	}
	return resultHandler.ServeHTTP
}

func (h *BaseHandler) getListMethods() []string {
	var useMethods []string
	for key := range h.handlerMethods {
		useMethods = append(useMethods, key)
	}
	useMethods = append(useMethods, http.MethodOptions)
	return useMethods
}

func (h *BaseHandler) add(path string, handler echo.HandlerFunc, route *echo.Group) {
	for key := range h.handlerMethods {
		switch key {
		case GET:
			route.GET(path, handler)
			break
		case POST:
			route.POST(path, handler)
			break
		case PUT:
			route.PUT(path, handler)
			break
		case DELETE:
			route.DELETE(path, handler)
			break
		case OPTIONS:
			route.OPTIONS(path, handler)
			break
		}
	}
}

func (h *BaseHandler) Connect(route *echo.Group, path string) {
	h.add(path, h.applyMiddleware(h), route)
}

func (h *BaseHandler) ServeHTTP(ctx echo.Context) error {
	h.PrintRequest(ctx)
	ok := true
	var hndlr hf.HandlerFunc

	hndlr, ok = h.handlerMethods[ctx.Request().Method]
	if ok {
		return hndlr(ctx)
	} else {
		//h.Log(ctx).Errorf("Unexpected http method: %s", ctx.Method())
		ctx.Request().Header.Set("Allow", strings.Join(h.getListMethods(), ", "))
		ctx.Response().WriteHeader(http.StatusInternalServerError)
	}
	return nil
}
