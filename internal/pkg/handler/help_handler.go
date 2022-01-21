package handler

import (
	"github.com/gorilla/mux"
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

func (h *HelpHandlers) PrintRequest(w http.ResponseWriter, r *http.Request) {
	//h.Log(w, r).Infof("Request: %s. From URL: %s", w, r.Method(), string(w, r.URI().Host())+string(w, r.Path()))
}

// GetInt64FromParam HTTPErrors
//		Status 400 handler_errors.InvalidParameters
func (h *HelpHandlers) GetInt64FromParam(w http.ResponseWriter, r *http.Request, name string) (int64, int, error) {
	vars := mux.Vars(r)
	number := vars[name]
	numberInt, err := strconv.ParseInt(number, 10, 64)
	if number == "" || err != nil {
		//h.Log(w, r).Infof("can't get parametrs %s, was got %v)", name, number)
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
func (h *HelpHandlers) GetPaginationFromQuery(w http.ResponseWriter, r *http.Request) (*Pagination, int, error) {
	limit, code, err := h.GetInt64FromQueries(w, r, "limit")
	if err != nil {
		return nil, code, err
	}

	if limit == EmptyQuery {
		limit = DefaultLimit
	}

	desc := h.GetBoolFromQueries(w, r, "desc")

	since, info := h.GetStringFromQueries(w, r, "since")
	if info == EmptyQuery {
		since = ""
	}
	return &Pagination{Since: since, Desc: desc, Limit: limit}, app.InvalidInt, nil
}

// GetInt64FromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetInt64FromQueries(w http.ResponseWriter, r *http.Request, name string) (int64, int, error) {
	number := r.URL.Query().Get(name)
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
func (h *HelpHandlers) GetBoolFromQueries(w http.ResponseWriter, r *http.Request, name string) bool {
	number := r.URL.Query().Get(name)
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
func (h *HelpHandlers) GetStringFromQueries(w http.ResponseWriter, r *http.Request, name string) (string, int) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return "", EmptyQuery
	}

	return value, app.InvalidInt
}

// GetStringFromParam HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetStringFromParam(w http.ResponseWriter, r *http.Request, name string) (string, int) {
	vars := mux.Vars(r)
	value := vars[name]
	if value == "" {
		return "", EmptyQuery
	}

	return value, app.InvalidInt
}

// GetArrayStringFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetArrayStringFromQueries(w http.ResponseWriter, r *http.Request, name string) ([]string, int) {
	values := r.URL.Query().Get(name)
	if values == "" {
		return nil, EmptyQuery
	}

	return strings.Split(values, ","), app.InvalidInt
}

// GetThreadSlugFromParam HTTPErrors
//		Status 400 handler_errors.InvalidParameters
func (h *HelpHandlers) GetThreadSlugFromParam(w http.ResponseWriter, r *http.Request, name string) (*thread.ThreadPK, int) {
	vars := mux.Vars(r)
	value := vars[name]
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

func (h *HelpHandlers) GetRequestBody(w http.ResponseWriter, r *http.Request, reqStruct easyjson.Unmarshaler) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	if err := easyjson.UnmarshalFromReader(r.Body, reqStruct); err != nil {
		return err
	}
	return nil
}
