package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
	"strconv"
	"strings"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"tech-db-forum/internal/pkg/utilits/delivery"
)

type Pagination struct {
	Limit int64
	Desc  bool
	Since string
}

const (
	EmptyQuery   = -2
	DefaultLimit = 100
)

type HelpHandlers struct {
	delivery.ErrorConvertor
}

func (h *HelpHandlers) PrintRequest(ctx echo.Context) {
	//h.Log(ctx).Infof("Request: %s. From URL: %s", ctx.Method(), string(ctx.URI().Host())+string(ctx.Path()))
}

// GetInt64FromParam HTTPErrors
//		Status 400 handler_errors.InvalidParameters
func (h *HelpHandlers) GetInt64FromParam(ctx echo.Context, name string) (int64, int, error) {
	number := ctx.Param(name)
	numberInt, err := strconv.ParseInt(number, 10, 64)
	if number == "" || err != nil {
		//h.Log(ctx).Infof("can't get parametrs %s, was got %v)", name, number)
		return app.InvalidInt, http.StatusBadRequest, handler_errors.InvalidParameters
	}
	return numberInt, app.InvalidInt, nil
}

// GetPaginationFromQuery Expected api param:
// 	Default value for limit - 100
//	Param since query any false "start number of values"
// 	Param limit query uint64 false "number values to return"
//	Param desc  query bool false "
// Errors:
// 	Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetPaginationFromQuery(ctx echo.Context) (*Pagination, int, error) {
	limit, code, err := h.GetInt64FromQueries(ctx, "limit")
	if err != nil {
		return nil, code, err
	}

	if limit == EmptyQuery {
		limit = DefaultLimit
	}

	desc := h.GetBoolFromQueries(ctx, "desc")

	since, info := h.GetStringFromQueries(ctx, "since")
	if info == EmptyQuery {
		since = ""
	}
	return &Pagination{Since: since, Desc: desc, Limit: limit}, app.InvalidInt, nil
}

// GetInt64FromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetInt64FromQueries(ctx echo.Context, name string) (int64, int, error) {
	number := ctx.QueryParam(name)
	if number == "" {
		return EmptyQuery, app.InvalidInt, nil
	}

	numberInt, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return app.InvalidInt, http.StatusBadRequest, handler_errors.InvalidQueries
	}

	return numberInt, app.InvalidInt, nil
}

// GetBoolFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetBoolFromQueries(ctx echo.Context, name string) bool {
	number := ctx.QueryParam(name)
	if number == "" {
		return false
	}

	numberInt, err := strconv.ParseBool(number)
	if err != nil {
		return false
	}

	return numberInt
}

// GetStringFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetStringFromQueries(ctx echo.Context, name string) (string, int) {
	value := ctx.QueryParam(name)
	if value == "" {
		return "", EmptyQuery
	}

	return value, app.InvalidInt
}

// GetStringFromParam HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetStringFromParam(ctx echo.Context, name string) (string, int) {
	value := ctx.Param(name)
	if value == "" {
		return "", EmptyQuery
	}

	return value, app.InvalidInt
}

// GetArrayStringFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetArrayStringFromQueries(ctx echo.Context, name string) ([]string, int) {
	values := ctx.QueryParam(name)
	if values == "" {
		return nil, EmptyQuery
	}

	return strings.Split(values, ","), app.InvalidInt
}

// GetThreadSlugFromParam HTTPErrors
//		Status 400 handler_errors.InvalidParameters
func (h *HelpHandlers) GetThreadSlugFromParam(ctx echo.Context, name string) (*thread.ThreadPK, int) {
	value := ctx.Param(name)
	if value == "" {
		return nil, EmptyQuery
	}

	res := &thread.ThreadPK{}
	res.SetSlug(value)
	id, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		res.SetId(id)
	}
	return res, app.InvalidInt
}

func (h *HelpHandlers) GetRequestBody(ctx echo.Context, reqStruct easyjson.Unmarshaler) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(ctx.Request().Body)

	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, reqStruct); err != nil {
		return err
	}
	return nil
}
