package glogger

import (
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *LogConfig) (Logger, error) {
	l, err := RegisterLog(*lc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return l, nil
}
