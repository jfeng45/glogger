package logrus

import (
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/config"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type LogrusFactory struct{}

// build logrus logger
func (mf *LogrusFactory) Build(lc *config.Logging) (glogger.Logger, error) {
	l, err := RegisterLogrusLog(lc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return l, nil
}
