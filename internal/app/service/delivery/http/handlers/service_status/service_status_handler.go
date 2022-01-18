package service_status_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/service/delivery/http"
	"tech-db-forum/internal/app/service/repository"
	bh "tech-db-forum/internal/pkg/handler"
)

type UserProfileHandler struct {
	serviceRepository repository.Repository
	bh.BaseHandler
}

func NewUserProfileHandler(log *logrus.Logger, rep repository.Repository) *UserProfileHandler {
	h := &UserProfileHandler{
		BaseHandler:       *bh.NewBaseHandler(log),
		serviceRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *UserProfileHandler) GET(ctx *routing.Context) error {
	stat, err := h.serviceRepository.GetStat()
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	h.Log(ctx).Debugf("get status of server %v", stat)
	h.Respond(ctx, http.StatusOK, http_delivery.ToStatusResponse(stat))
	return nil
}
