package mlog

var (
	logger *Logger
)

func Log() *Logger {
	if logger == nil {
		cnf, err := loadConfig()
		if err != nil {
			panic(err)
		}
		logger = &Logger{
			config: cnf,
		}
	}
	return logger
}
