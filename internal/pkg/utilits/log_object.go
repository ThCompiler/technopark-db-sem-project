package utilits

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LogObject struct {
	log *logrus.Logger
}

func NewLogObject(log *logrus.Logger) LogObject {
	return LogObject{log: log}
}

func (l *LogObject) BaseLog() *logrus.Logger {
	return l.log
}

func (l *LogObject) Log(ctx echo.Context) *logrus.Entry {
	if ctx == nil {
		return l.log.WithField("type", "base_log")
	}
	ctxLogger := ctx.Request().Context().Value("logger")
	logger := l.log.WithField("urls", ctx.Request().URL)
	if ctxLogger != nil {
		if log, ok := ctxLogger.(*logrus.Entry); ok {
			logger = log
		}
	}
	return logger
}
