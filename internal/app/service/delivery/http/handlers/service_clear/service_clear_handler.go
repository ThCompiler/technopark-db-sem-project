package service_clear_handler

import (
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

func (h *ServiceClearHandler) POST(w http.ResponseWriter, r *http.Request) {
	err := h.serviceRepository.Clear()
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debug("bd was cleared")
	w.WriteHeader(http.StatusOK)
	return
}
