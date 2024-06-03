package gins

import (
	"github.com/jiftle/mlog/glog"
)

// Log returns an instance of glog.Logger.
// The parameter <name> is the name for the instance.
func Log(name ...string) *glog.Logger {
	instanceName := glog.DefaultName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	logger := glog.Instance(instanceName)

	m, err := loadConfig()
	if err != nil {
		panic(err)
	}
	if err := logger.SetConfigWithMap(m); err != nil {
		panic(err)
	}
	return logger
}
