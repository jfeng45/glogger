package glogger

import (
	"github.com/pkg/errors"
)

// receiver for logrus factory
type LogrusFactory struct{}

// build logrus logger
func (mf *LogrusFactory) Build(lc *LogConfig) (Logger, error) {
	l, err := RegisterLogrusLog(*lc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return l, nil
}
