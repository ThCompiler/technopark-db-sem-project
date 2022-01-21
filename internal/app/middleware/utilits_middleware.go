package middleware

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	hf "tech-db-forum/internal/pkg/handler/handler_interfaces"
	"tech-db-forum/internal/pkg/utilits"
)

type UtilitiesMiddleware struct {
	log utilits.LogObject
}

func NewUtilitiesMiddleware(log *logrus.Logger) UtilitiesMiddleware {
	return UtilitiesMiddleware{
		log: utilits.NewLogObject(log),
	}
}
/*
func (mw *UtilitiesMiddleware) CheckPanic() hf.Handler {
	return hf.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(log *logrus.Entry, w http.ResponseWriter, r *http.Request) {
			if err := recover(); err != nil {
				responseErr := http.StatusInternalServerError

				log.Errorf("detacted critical error: %v, with stack: %s", err, debug.Stack())
				w.WriteHeader(responseErr)
			}
		}(mw.log.Log(w, r), w, r)
		return w, r.Next()
	})
}

func (mw *UtilitiesMiddleware) UpgradeLogger() hf.Handler {
	return hf.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*start := time.Now()
		upgradeLogger := mw.log.BaseLog().WithFields(logrus.Fields{
			"urls":        w, r.URI().String(),
			"method":      string(w, r.Method()),
			"remote_addr": w, r.RemoteAddr(),
			"work_time":   time.Since(start).Milliseconds(),
			"req_id":      uuid.NewV4(),
		})
		w, r.SetUserValue("logger", upgradeLogger)
		upgradeLogger.Info("Log was upgraded")
*/
		err := w, r.Next()

		/*executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
		return err*/
		return err
	})
}*/
