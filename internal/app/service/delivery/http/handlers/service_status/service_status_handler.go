package service_status_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/service/delivery/http"
	"tech-db-forum/internal/app/service/repository"
	bh "tech-db-forum/internal/pkg/handler"
)

type ServiceStatusHandler struct {
	serviceRepository repository.Repository
	bh.BaseHandler
}

func NewServiceStatusHandler(log *logrus.Logger, rep repository.Repository) *ServiceStatusHandler {
	h := &ServiceStatusHandler{
		BaseHandler:       *bh.NewBaseHandler(log),
		serviceRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ServiceStatusHandler) GET(ctx *routing.Context) error {
	stat, err := h.serviceRepository.GetStat()
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get status of server %v", stat)
	h.Respond(ctx, http.StatusOK, http_delivery.ToStatusResponse(stat))
	return nil
}
