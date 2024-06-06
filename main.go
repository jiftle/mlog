package main

import (
	"github.com/jiftle/mlog/g"
	"github.com/jiftle/mlog/mlog"
)

func main() {
	g.Log().Debug("this is debug level info !")
	mlog.Log().Debugf("debug level: %d %d %s", 10, 20, "世界和平")

	mlog.Log().Info("this is debug info info !")
	mlog.Log().Infof("info level: %d %d %s", 10, 20, "世界和平")

	mlog.Log().Warning("this is warn info !")
	mlog.Log().Warningf("warn level: %d %d %s", 10, 20, "世界和平")

	mlog.Log().Error("this is error info !")
	mlog.Log().Errorf("error level: %d %d %s", 10, 20, "世界和平")
}
