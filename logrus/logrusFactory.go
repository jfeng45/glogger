package logrus

import (
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/logconfig"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type LogrusFactory struct{}

// build logrus logger
func (mf *LogrusFactory) Build(lc *logconfig.LogConfig) (glogger.Logger, error) {
	l, err := RegisterLogrusLog(*lc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return l, nil
}
