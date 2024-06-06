package mlog

type Logger struct {
	config *Config
}

func (s *Logger) Debug(v ...interface{}) {
	if s.checkLevel(LEVEL_DEBU) {
		level := "DEBU"
		format := ""
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Debugf(format string, v ...interface{}) {
	if s.checkLevel(LEVEL_DEBU) {
		level := "DEBU"
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Info(v ...interface{}) {
	if s.checkLevel(LEVEL_INFO) {
		level := "INFO"
		format := ""
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Infof(format string, v ...interface{}) {
	if s.checkLevel(LEVEL_INFO) {
		level := "INFO"
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Warn(v ...interface{}) {
	if s.checkLevel(LEVEL_WARN) {
		level := "WARN"
		format := ""
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Warnf(format string, v ...interface{}) {
	if s.checkLevel(LEVEL_WARN) {
		level := "WARN"
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Warning(v ...interface{}) {
	s.Warn(v...)
}

func (s *Logger) Warningf(format string, v ...interface{}) {
	s.Warnf(format, v...)
}

func (s *Logger) Error(v ...interface{}) {
	if s.checkLevel(LEVEL_ERRO) {
		level := "ERRO"
		format := ""
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}

func (s *Logger) Errorf(format string, v ...interface{}) {
	if s.checkLevel(LEVEL_ERRO) {
		level := "ERRO"
		if s.config.Stdout {
			s.printToStdout(level, format, v...)
		}
		s.printToFile(level, format, v...)
	}
}
