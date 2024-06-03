package g

import (
	"github.com/jiftle/mlog/frame/gins"
	"github.com/jiftle/mlog/glog"
)

// Log returns an instance of glog.Logger.
// The parameter <name> is the name for the instance.
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}
