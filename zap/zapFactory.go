package zap

import (
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/config"
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *config.Logging) (glogger.Logger, error) {
	l, err := RegisterLog(lc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return l, nil
}
