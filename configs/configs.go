package configs

var (
	logger *Logger
)

func GetLogger() *Logger {
	if logger == nil {
		logger = NewLogger()
	}

	return logger
}