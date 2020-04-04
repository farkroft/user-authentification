package log

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/auth-service/external/util"
)

// WIBFormatter struct
type WIBFormatter struct {
	logrus.Formatter
}

// NewLogger format logrus
func NewLogger() {
	logrus.SetLevel(logrus.InfoLevel)

	switch logrus.GetLevel() {
	case logrus.ErrorLevel:
		logrus.SetReportCaller(true)
		logrus.SetFormatter(WIBFormatter{&logrus.JSONFormatter{}})
	case logrus.InfoLevel:
		logrus.SetFormatter(WIBFormatter{&logrus.TextFormatter{FullTimestamp: true}})
	}
}

// Format return time formatted with timezone
func (w WIBFormatter) Format(e *logrus.Entry) ([]byte, error) {
	value := util.WIBTimezone(e.Time.UTC())
	e.Time = value
	return w.Formatter.Format(e)
}

// Errorf error with format string
func Errorf(str string, err error) {
	s := str + " %s"
	logrus.Errorf(s, err)
}

// Infoln log info message
func Infoln(str string) {
	logrus.Infoln(str)
}
