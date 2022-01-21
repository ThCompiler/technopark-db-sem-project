package service_clear_handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/service/repository"
	bh "tech-db-forum/internal/pkg/handler"
)

type ServiceClearHandler struct {
	serviceRepository repository.Repository
	bh.BaseHandler
}

func NewServiceClearHandler(log *logrus.Logger, rep repository.Repository) *ServiceClearHandler {
	h := &ServiceClearHandler{
		BaseHandler:       *bh.NewBaseHandler(log),
		serviceRepository: rep,
	}
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *ServiceClearHandler) POST(ctx echo.Context) error {
	err := h.serviceRepository.Clear()
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debug("bd was cleared")
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}
