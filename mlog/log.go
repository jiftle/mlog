package mlog

var (
	logger *Logger
)

func Log() *Logger {
	if logger == nil {
		cnf, err := loadConfig()
		if err != nil {
			logger = &Logger{
				config: &Config{
					Path:   "logs",
					Level:  "debug",
					Stdout: true,
				},
			}
			logger.SetLevelStr("debug")
		} else {
			logger = &Logger{
				config: cnf,
			}
			logger.SetLevelStr(cnf.Level)
		}
	}
	return logger
}
